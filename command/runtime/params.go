package runtime

const (
	attackerFlag          = "attacker"
	maxValidatorIndexFlag = "max-validator-index"
	strategyFlag          = "strategy"
	listFlag              = "list"
)

type updateParam struct {
	attacker          string
	maxValidatorIndex int
	strategy          string
	listLibrary       bool
}

var (
	params = &updateParam{
		attacker:          "",
		strategy:          "",
		maxValidatorIndex: -1,
		listLibrary:       false,
	}
)
