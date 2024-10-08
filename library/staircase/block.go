package staircase

import (
	"fmt"
	"github.com/tsinghua-cel/strategy-gen/types"
)

// 1. 每个epoch 的区块不出块，投票不广播
// 2. 每个epoch的最后一个作恶者出块，打包所有的 投票，并且延迟广播。
// 3.设block slot为t , receive block: (32 - t mod 32)*12s , broadcast block：在上一个的基础上12*12s

func GenSlotStrategy(latestHackDutySlot int, epoch int64, cas int) []types.SlotStrategy {
	strategys := make([]types.SlotStrategy, 0)
	switch cas {
	case 0:
		strategy := types.SlotStrategy{
			Slot:    "every",
			Level:   0,
			Actions: make(map[string]string),
		}
		strategy.Actions["BlockBeforeSign"] = "return"
		strategy.Actions["AttestBeforeSign"] = fmt.Sprintf("return")
		strategys = append(strategys, strategy)

	case 1:
		{
			strategy := types.SlotStrategy{
				Slot:    "every",
				Level:   0,
				Actions: make(map[string]string),
			}
			strategy.Actions["BlockBeforeSign"] = "return"
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
			stageI := (32 - latestHackDutySlot%32) * 12
			stageII := 12 * 12

			strategy.Actions["AttestBeforeSign"] = fmt.Sprintf("return")

			strategy.Actions["BlockBeforeSign"] = "packPooledAttest"
			strategy.Actions["BlockDelayForReceiveBlock"] = fmt.Sprintf("%s:%d", "delayWithSecond", stageI)
			strategy.Actions["BlockBeforeBroadCast"] = fmt.Sprintf("%s:%d", "delayWithSecond", stageII)

			strategys = append(strategys, strategy)
		}

	}
	return strategys

}
