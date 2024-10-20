package library

import (
	aiattack "github.com/tsinghua-cel/strategy-gen/library/ai"
	"github.com/tsinghua-cel/strategy-gen/library/confuse"
	"github.com/tsinghua-cel/strategy-gen/library/exante"
	"github.com/tsinghua-cel/strategy-gen/library/five"
	"github.com/tsinghua-cel/strategy-gen/library/four"
	"github.com/tsinghua-cel/strategy-gen/library/one"
	"github.com/tsinghua-cel/strategy-gen/library/sandwich"
	"github.com/tsinghua-cel/strategy-gen/library/staircase"
	"github.com/tsinghua-cel/strategy-gen/library/three"
	"github.com/tsinghua-cel/strategy-gen/library/two"
	"github.com/tsinghua-cel/strategy-gen/library/unrealized"
	"github.com/tsinghua-cel/strategy-gen/library/withholding"
	"github.com/tsinghua-cel/strategy-gen/types"
	"sync"
)

type Strategy interface {
	Name() string
	Run(param types.LibraryParams, feedbacker types.FeedBacker)
	Description() string
}

var (
	allStrategies = sync.Map{}
)

func Init() {
	register(&one.Instance{})
	register(&two.Instance{})
	register(&three.Instance{})
	register(&four.Instance{})
	register(&five.Instance{})
	register(&exante.Instance{})
	register(&sandwich.Instance{})
	register(&withholding.Instance{})
	register(&unrealized.Instance{})
	register(&staircase.Instance{})
	register(&confuse.Instance{})
	register(&aiattack.Instance{})
}

func register(ins Strategy) {
	allStrategies.Store(ins.Name(), ins)
}

func GetStrategy(name string) (Strategy, bool) {
	if v, ok := allStrategies.Load(name); ok {
		return v.(Strategy), true
	}
	return nil, false
}

func GetAllStrategies() map[string]Strategy {
	strategies := make(map[string]Strategy)
	allStrategies.Range(func(k, v interface{}) bool {
		strategies[k.(string)] = v.(Strategy)
		return true
	})
	return strategies
}
