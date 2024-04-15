package actionset

type AType int

const (
	AnyAction AType = iota
	BlockAction
	AttestAction
)

type Action interface {
	Name() string
	MaxParam() int
	MinParam() int
	ActionType() AType
	DefaultParam() []interface{}
	RandomParam() []interface{}
	Desc() string
}

var anyAction = []Action{
	NullAction{},
	ReturnAction{},
	ContinueAction{},
	AbortAction{},
	SkipAction{},
	ExitAction{},
	delayWithSecondAction{},
	delayToNextSlotAction{},
	delayToAfterNextSlotAction{},
	delayToNextNEpochStartAction{},
	delayToNextNEpochEndAction{},
	delayToNextNEpochHalfAction{},
	delayToEpochEndAction{},
	delayHalfEpochAction{},
}

var attestAction = []Action{
	storeSignedAttestAction{},
	rePackAttestationAction{},
}
var blockAction = []Action{}

func GetBlockActionSet() []Action {
	a := make([]Action, 0)
	a = append(a, anyAction...)
	a = append(a, blockAction...)

	return a
}

func GetAttestActionSet() []Action {
	a := make([]Action, 0)
	a = append(a, anyAction...)
	a = append(a, attestAction...)
	return a
}

func GetAllActionSet() []Action {
	all := make([]Action, 0)
	all = append(all, anyAction...)
	all = append(all, blockAction...)
	all = append(all, attestAction...)
	return all
}
