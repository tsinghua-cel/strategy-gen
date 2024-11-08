package four

import (
	"fmt"
	"github.com/tsinghua-cel/strategy-gen/types"
	"strconv"
)

func getSlotStrategy(epoch int64, slot string, isLatestHackDuty bool) types.SlotStrategy {
	strategy := types.SlotStrategy{
		Slot:    slot,
		Level:   0,
		Actions: make(map[string]string),
	}
	switch epoch % 3 {
	case 0, 1:
		strategy.Actions["BlockBeforeSign"] = "return"
		strategy.Actions["AttestBeforeSign"] = fmt.Sprintf("return")
		return strategy

	case 2:
		if isLatestHackDuty {
			strategy.Level = 1
			islot, _ := strconv.Atoi(slot)
			stageI := (32 - islot%32) * 12
			stageII := 12 * 12

			strategy.Actions["AttestBeforeSign"] = fmt.Sprintf("return")

			strategy.Actions["BlockBeforeSign"] = "packPooledAttest"
			strategy.Actions["BlockDelayForReceiveBlock"] = fmt.Sprintf("%s:%d", "delayWithSecond", stageI)
			strategy.Actions["BlockBeforeBroadCast"] = fmt.Sprintf("%s:%d", "delayWithSecond", stageII)

		} else {
			strategy.Actions["BlockBeforeSign"] = "return"
			strategy.Actions["AttestAfterSign"] = fmt.Sprintf("addAttestToPool")
			strategy.Actions["AttestBeforeBroadCast"] = fmt.Sprintf("return")
		}
	}
	return strategy
}

func GenSlotStrategy(allDuties []types.ProposerDuty, epoch int64) []types.SlotStrategy {
	strategys := make([]types.SlotStrategy, 0)
	for i := 0; i < len(allDuties); i++ {
		s := getSlotStrategy(epoch, allDuties[i].Slot, i == len(allDuties)-1)
		strategys = append(strategys, s)
	}
	return strategys
}
