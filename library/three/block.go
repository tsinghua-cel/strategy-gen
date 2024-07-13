package three

import (
	"fmt"
	"github.com/tsinghua-cel/strategy-gen/types"
)

func BlockStrategyForEpoch0(actions map[string]string) {
	// nothing to do
}

func AttestStrategyForEpoch0(actions map[string]string) {
	actions["AttestAfterSign"] = fmt.Sprintf("addAttestToPool")
	actions["AttestBeforeBroadCast"] = fmt.Sprintf("return")
}

func BlockStrategyForEpoch1(actions map[string]string) {
	actions["BlockBeforeSign"] = "packPooledAttest"
	// block delay to next epoch-end slot
	actions["BlockBeforeBroadCast"] = fmt.Sprintf("delayToNextNEpochEnd:%d", 1)
}

func AttestStrategyForEpoch1(actions map[string]string) {
	actions["AttestAfterSign"] = fmt.Sprintf("addAttestToPool")
	actions["AttestBeforeBroadCast"] = fmt.Sprintf("return")
}

func GenSlotStrategy(latestHackDutySlot int, epoch int64) []types.SlotStrategy {
	strategys := make([]types.SlotStrategy, 0)
	switch epoch % 2 {
	case 0:
		strategy := types.SlotStrategy{
			Slot:    "every",
			Level:   0,
			Actions: make(map[string]string),
		}
		BlockStrategyForEpoch0(strategy.Actions)
		AttestStrategyForEpoch0(strategy.Actions)
		strategys = append(strategys, strategy)

	case 1:
		{
			strategy := types.SlotStrategy{
				Slot:    "every",
				Level:   0,
				Actions: make(map[string]string),
			}
			BlockStrategyForEpoch0(strategy.Actions)
			AttestStrategyForEpoch0(strategy.Actions)
			strategys = append(strategys, strategy)
		}
		{
			strategy := types.SlotStrategy{
				Slot:    fmt.Sprintf("%d", latestHackDutySlot),
				Level:   1,
				Actions: make(map[string]string),
			}
			BlockStrategyForEpoch1(strategy.Actions)
			AttestStrategyForEpoch1(strategy.Actions)
			strategys = append(strategys, strategy)
		}

	}
	return strategys

}
