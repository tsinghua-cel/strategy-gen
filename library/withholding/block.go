package withholding

import (
	"fmt"
	"github.com/tsinghua-cel/strategy-gen/pointset"
	"github.com/tsinghua-cel/strategy-gen/types"
	"github.com/tsinghua-cel/strategy-gen/utils"
	"strconv"
)

func BlockStrategy(cur, end int, actions map[string]string) {
	endStage := end + 1 - cur
	endStage += 32 / 2 // add half epoch.
	point := pointset.GetPointByName("BlockBeforeBroadCast")
	actions[point] = fmt.Sprintf("%s:%d", "delayWithSecond", endStage*12)
}

func GenSlotStrategy(allHacks []interface{}) []types.SlotStrategy {
	if len(allHacks) == 0 {
		return nil
	}
	strategys := make([]types.SlotStrategy, 0)

	// only use the last subduties
	subduties := allHacks[len(allHacks)-1]

	duties := subduties.([]utils.ProposerDuty)
	end, _ := strconv.Atoi(duties[len(duties)-1].Slot)

	for i := 0; i < len(duties); i++ {
		slot, _ := strconv.Atoi(duties[i].Slot)
		strategy := types.SlotStrategy{
			Slot:    duties[i].Slot,
			Level:   1,
			Actions: make(map[string]string),
		}
		BlockStrategy(slot, end, strategy.Actions)
		strategys = append(strategys, strategy)
	}

	return strategys

}
