package runtime

const (
	attackerFlag          = "attacker"
	maxValidatorIndexFlag = "max-validator-index"
	minValidatorIndexFlag = "min-validator-index"
	strategyFlag          = "strategy"
	listFlag              = "list"
	logFlag               = "log"
)

type updateParam struct {
	attacker          string
	maxValidatorIndex int
	minValidatorIndex int
	strategy          string
	listLibrary       bool
	logPath           string
}

var (
	params = &updateParam{
		attacker:          "",
		strategy:          "",
		logPath:           "",
		maxValidatorIndex: -1,
		minValidatorIndex: 0,
		listLibrary:       false,
	}
)
