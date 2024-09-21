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
	epochPerStage  = 8 // every 8 epoch is a stage, to change strategy.
	stageCache, _  = lru.New(100)
)

type stageInfo struct {
	duties  []utils.ProposerDuty
	endSlot int64
}

func BlockStrategy(cur, end int, actions map[string]string) {
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

func GenSlotStrategy(allHacks []interface{}) []types.SlotStrategy {
	if len(allHacks) == 0 {
		return nil
	}
	strategys := make([]types.SlotStrategy, 0)
	duties := make([]utils.ProposerDuty, 0)
	latestDuty := allHacks[len(allHacks)-1].(utils.ProposerDuty)
	endSlot, _ := strconv.ParseInt(latestDuty.Slot, 10, 64)
	for _, iduty := range allHacks {
		duty := iduty.(utils.ProposerDuty)
		duties = append(duties, duty)
		slot, _ := strconv.ParseInt(duty.Slot, 10, 64)
		strategy := types.SlotStrategy{
			Slot:    duty.Slot,
			Level:   1,
			Actions: make(map[string]string),
		}
		BlockStrategy(int(slot), int(endSlot), strategy.Actions)
		strategys = append(strategys, strategy)
	}
	epoch := utils.SlotTool{32}.SlotToEpoch(endSlot)
	stageCache.Add(epoch, stageInfo{duties: duties, endSlot: endSlot})

	return strategys

}
