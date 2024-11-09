package randomdelay

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/tsinghua-cel/strategy-gen/types"
	"math/rand"
	"strconv"
	"time"
)

type ActionPoint string

const (
	AttestBeforeBroadCast     ActionPoint = "AttestBeforeBroadCast"
	AttestAfterBroadCast      ActionPoint = "AttestAfterBroadCast"
	AttestBeforeSign          ActionPoint = "AttestBeforeSign"
	AttestAfterSign           ActionPoint = "AttestAfterSign"
	AttestBeforePropose       ActionPoint = "AttestBeforePropose"
	AttestAfterPropose        ActionPoint = "AttestAfterPropose"
	BlockDelayForReceiveBlock ActionPoint = "BlockDelayForReceiveBlock"
	BlockBeforeBroadCast      ActionPoint = "BlockBeforeBroadCast"
	BlockAfterBroadCast       ActionPoint = "BlockAfterBroadCast"
	BlockBeforeSign           ActionPoint = "BlockBeforeSign"
	BlockAfterSign            ActionPoint = "BlockAfterSign"
	BlockBeforePropose        ActionPoint = "BlockBeforePropose"
	BlockAfterPropose         ActionPoint = "BlockAfterPropose"
	BlockGetNewParentRoot     ActionPoint = "BlockGetNewParentRoot"
)

var (
	stageCache, _ = lru.New(100)
	allPoints     = []ActionPoint{
		AttestBeforeBroadCast,
		AttestAfterBroadCast,
		AttestBeforeSign,
		AttestAfterSign,
		AttestBeforePropose,
		AttestAfterPropose,
		BlockDelayForReceiveBlock,
		BlockBeforeBroadCast,
		BlockAfterBroadCast,
		BlockBeforeSign,
		BlockAfterSign,
		BlockBeforePropose,
		BlockAfterPropose,
		BlockGetNewParentRoot,
	}
	beforePoint = []ActionPoint{
		AttestBeforeBroadCast,
		//AttestAfterBroadCast,
		AttestBeforeSign,
		//AttestAfterSign,
		AttestBeforePropose,
		//AttestAfterPropose,
		BlockDelayForReceiveBlock,
		BlockBeforeBroadCast,
		//BlockAfterBroadCast,
		BlockBeforeSign,
		//BlockAfterSign,
		BlockBeforePropose,
		//BlockAfterPropose,
		BlockGetNewParentRoot,
	}
)

type RSet struct {
	items []ActionPoint
}

func (s *RSet) Add(item []ActionPoint) {
	s.items = append(s.items, item...)
}

func (s *RSet) PopRandom() (ActionPoint, error) {
	if len(s.items) == 0 {
		return ActionPoint(""), fmt.Errorf("set is empty")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(s.items))
	item := s.items[index]
	s.items = append(s.items[:index], s.items[index+1:]...)

	return item, nil
}

func strategy2(idx, cur, end int, actions map[string]string) {
	// delay random time before all broadcast point.
	baseDelay := 30 + idx*5
	for _, point := range beforePoint {
		delay := rand.Intn(20) + baseDelay
		actions[string(point)] = fmt.Sprintf("%s:%d", "delayWithSecond", delay)
	}
}

func strategy1(idx, cur, end int, actions map[string]string) {
	// random select some point and delay random time.
	baseDelay := 30 + idx*5
	basePointCount := 4

	all := &RSet{
		items: make([]ActionPoint, 0),
	}
	all.Add(allPoints)

	count := rand.Intn(3) + basePointCount
	for i := 0; i < count; i++ {
		if point, err := all.PopRandom(); err != nil {
			break
		} else {
			delay := rand.Intn(20) + baseDelay
			actions[string(point)] = fmt.Sprintf("%s:%d", "delayWithSecond", delay)
		}
	}
}

func BlockStrategy(idx, cur, end int, actions map[string]string) {
	if cur%300 < 150 {
		strategy2(idx, cur, end, actions)
	} else {
		strategy1(idx, cur, end, actions)
	}
}

func GenSlotStrategy(allHacks []types.ProposerDuty) []types.SlotStrategy {
	if len(allHacks) == 0 {
		return nil
	}
	latestDuty := allHacks[len(allHacks)-1]
	endSlot, _ := strconv.ParseInt(latestDuty.Slot, 10, 64)
	strategys := make([]types.SlotStrategy, 0)
	for i, duty := range allHacks {
		slot, _ := strconv.ParseInt(duty.Slot, 10, 64)
		strategy := types.SlotStrategy{
			Slot:    duty.Slot,
			Level:   1,
			Actions: make(map[string]string),
		}
		BlockStrategy(i, int(slot), int(endSlot), strategy.Actions)
		strategys = append(strategys, strategy)
	}

	return strategys

}
