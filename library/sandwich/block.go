package sandwich

import (
	"fmt"
	"github.com/tsinghua-cel/strategy-gen/types"
	"github.com/tsinghua-cel/strategy-gen/utils"
)

func GenSlotStrategy(duties []interface{}) []types.SlotStrategy {
	strategys := make([]types.SlotStrategy, 0)
	for i := 0; i < len(duties); i++ {
		duty := duties[i].([]utils.ProposerDuty)
		if len(duty) != 3 {
			continue
		}
		a := duty[0]
		//b := duty[1]
		c := duty[2]

		slotStrategy := types.SlotStrategy{
			Slot:    fmt.Sprintf("%s", c.Slot),
			Level:   1,
			Actions: make(map[string]string),
		}
		slotStrategy.Actions["BlockGetNewParentRoot"] = fmt.Sprintf("modifyParentRoot:%s", a.Slot)
		// todo: Is it need delay for block or attest ?
		strategys = append(strategys, slotStrategy)
	}

	return strategys

}
