package confuse

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/tsinghua-cel/strategy-gen/types"
	"github.com/tsinghua-cel/strategy-gen/utils"
	"strconv"
)

var (
	SecondsPerSlot = 12
	epochPerStage  = int64(8) // every 8 epoch is a stage, to change strategy.
	stageCache, _  = lru.New(100)
)

type stageInfo struct {
	duties  []utils.ProposerDuty
	endSlot int64
}

func BlockStrategy(idx, cur, end int, actions map[string]string) {
	if cur != end {
		// not propose block.
		// add attest to pool, not broadcast.
		actions["BlockBeforeSign"] = "return"
		actions["AttestAfterSign"] = fmt.Sprintf("addAttestToPool")
		actions["AttestBeforeBroadCast"] = fmt.Sprintf("return")
	}
	if cur == end {
		// propose block, and pack all pooled attest.
		// modify parent to last epoch endslot.
		actions["BlockBeforeSign"] = "packPooledAttest"
		lastEpoch := utils.SlotTool{32}.SlotToEpoch(int64(end)) - 1
		if data, exist := stageCache.Get(lastEpoch); exist {
			info := data.(stageInfo)
			actions["BlockGetNewParentRoot"] = fmt.Sprintf("modifyParentRoot:%d", info.endSlot)
		}
	}
}

func BlockStrategy2(idx, cur, end int, actions map[string]string) {
	// all block broadcast delay to end slot.
	// a-b-c-d-e delay to e + 1, e + 2, e + 3, e + 4, e + 0.

	if cur != end {
		// block broadcast delay to end slot.
		delay := (end-cur)*SecondsPerSlot + idx + 1
		actions["BlockBeforeBroadCast"] = fmt.Sprintf("delayWithSecond:%d", delay)
		// attest broadcast delay to end slot.
		actions["AttestBeforeBroadCast"] = fmt.Sprintf("delayWithSecond:%d", delay)
	}
}

func BlockStrategy3(idx, cur, lastSlot int, actions map[string]string) {
	// modify parent root to lastSlot.
	actions["BlockGetNewParentRoot"] = fmt.Sprintf("modifyParentRoot:%d", lastSlot)
	// don't broadcast attest.
	actions["AttestBeforeBroadCast"] = fmt.Sprintf("return")
}

func GenSlotStrategy(allHacks []interface{}) []types.SlotStrategy {
	if len(allHacks) == 0 {
		return nil
	}
	strategys := make([]types.SlotStrategy, 0)
	duties := make([]utils.ProposerDuty, 0)
	latestDuty := allHacks[len(allHacks)-1].(utils.ProposerDuty)
	endSlot, _ := strconv.ParseInt(latestDuty.Slot, 10, 64)
	epoch := utils.SlotTool{32}.SlotToEpoch(endSlot)
	lastEpochInfo, haveLast := stageCache.Get(epoch - 1)
	for i, iduty := range allHacks {
		duty := iduty.(utils.ProposerDuty)
		duties = append(duties, duty)
		slot, _ := strconv.ParseInt(duty.Slot, 10, 64)
		strategy := types.SlotStrategy{
			Slot:    duty.Slot,
			Level:   1,
			Actions: make(map[string]string),
		}
		var lastSlot int64 = slot - 1 // normally last slot.
		if i == 0 && haveLast {
			lastSlot = lastEpochInfo.(stageInfo).endSlot
		} else if i > 0 {
			lastSlot, _ = strconv.ParseInt(allHacks[i-1].(utils.ProposerDuty).Slot, 10, 64)
		}

		switch epoch / epochPerStage {
		case 1:
			BlockStrategy(i, int(slot), int(endSlot), strategy.Actions)
		case 2:
			BlockStrategy2(i, int(slot), int(endSlot), strategy.Actions)
		case 3:
			BlockStrategy3(i, int(slot), int(lastSlot), strategy.Actions)
		default:
			BlockStrategy(i, int(slot), int(endSlot), strategy.Actions)
		}

		strategys = append(strategys, strategy)
	}
	stageCache.Add(epoch, stageInfo{duties: duties, endSlot: endSlot})

	return strategys

}
