package two

import (
	"fmt"
	"github.com/tsinghua-cel/strategy-gen/types"
)

func BlockStrategyForEpoch0(actions map[string]string) {
	// nothing to do
}

func AttestStrategyForEpoch0(actions map[string]string) {
	actions["AttestBeforeSign"] = fmt.Sprintf("return")
}

func BlockStrategyForEpoch2(actions map[string]string) {
	actions["BlockBeforeSign"] = "packPooledAttest"
	// block delay to next epoch-end slot
	actions["BlockBeforeBroadCast"] = fmt.Sprintf("delayToNextNEpochStart:%d", 2)
}

func AttestStrategyForEpoch2(actions map[string]string) {
	actions["AttestAfterSign"] = fmt.Sprintf("addAttestToPool")
	actions["AttestBeforeBroadCast"] = fmt.Sprintf("return")
}

func GenSlotStrategy(latestHackDutySlot int, epoch int64) []types.SlotStrategy {
	strategys := make([]types.SlotStrategy, 0)
	switch epoch % 3 {
	case 0, 1:
		strategy := types.SlotStrategy{
			Slot:    "every",
			Level:   0,
			Actions: make(map[string]string),
		}
		strategy.Actions["BlockBeforePropose"] = "return"
		strategy.Actions["AttestBeforeSign"] = fmt.Sprintf("return")
		strategys = append(strategys, strategy)

	case 2:
		{
			strategy := types.SlotStrategy{
				Slot:    "every",
				Level:   0,
				Actions: make(map[string]string),
			}
			strategy.Actions["BlockBeforePropose"] = "return"
			strategy.Actions["AttestAfterSign"] = fmt.Sprintf("addAttestToPool")
			strategy.Actions["AttestBeforeBroadCast"] = fmt.Sprintf("return")
			strategys = append(strategys, strategy)
		}
		{
			strategy := types.SlotStrategy{
				Slot:    fmt.Sprintf("%d", latestHackDutySlot),
				Level:   1,
				Actions: make(map[string]string),
			}
			strategy.Actions["AttestBeforeSign"] = fmt.Sprintf("return")

			strategy.Actions["BlockBeforeSign"] = "packPooledAttest"
			strategy.Actions["BlockDelayForReceiveBlock"] = fmt.Sprintf("delayToNextNEpochStart:%d", 1)
			// block delay to next epoch-end slot
			strategy.Actions["BlockBeforeBroadCast"] = fmt.Sprintf("delayToNextNEpochStart:%d", 1)

			strategys = append(strategys, strategy)
		}

	}
	return strategys

}
