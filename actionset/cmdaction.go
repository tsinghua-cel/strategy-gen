package actionset

var (
	_ Action = NullAction{}
	_ Action = ReturnAction{}
	_ Action = ContinueAction{}
	_ Action = AbortAction{}
	_ Action = SkipAction{}
	_ Action = ExitAction{}
)

type baseCmdAction struct {
}

type NullAction struct{}

func (n NullAction) DefaultParam() []interface{} { return []interface{}{} }

func (n NullAction) RandomParam() []interface{} { return []interface{}{} }

func (n NullAction) Desc() string { return "# Null action does nothing" }

func (n NullAction) Name() string { return "null" }

func (n NullAction) MaxParam() int { return 0 }

func (n NullAction) MinParam() int { return 0 }

func (n NullAction) ActionType() AType { return AnyAction }

type ReturnAction struct{}

func (r ReturnAction) DefaultParam() []interface{} { return []interface{}{} }

func (r ReturnAction) RandomParam() []interface{} { return []interface{}{} }

func (r ReturnAction) Desc() string { return "# Return action returns from the current function" }

func (r ReturnAction) Name() string { return "return" }

func (r ReturnAction) MaxParam() int { return 0 }

func (r ReturnAction) MinParam() int { return 0 }

func (r ReturnAction) ActionType() AType { return AnyAction }

type ContinueAction struct{}

func (c ContinueAction) DefaultParam() []interface{} { return []interface{}{} }

func (c ContinueAction) RandomParam() []interface{} { return []interface{}{} }

func (c ContinueAction) Desc() string { return "# Continue action continues to the next iteration" }

func (c ContinueAction) Name() string { return "continue" }

func (c ContinueAction) MaxParam() int { return 0 }

func (c ContinueAction) MinParam() int { return 0 }

func (c ContinueAction) ActionType() AType { return AnyAction }

type AbortAction struct{}

func (a AbortAction) DefaultParam() []interface{} { return []interface{}{} }

func (a AbortAction) RandomParam() []interface{} { return []interface{}{} }

func (a AbortAction) Desc() string { return "# Abort action aborts the current function" }

func (a AbortAction) Name() string { return "abort" }

func (a AbortAction) MaxParam() int { return 0 }

func (a AbortAction) MinParam() int { return 0 }

func (a AbortAction) ActionType() AType { return AnyAction }

type SkipAction struct{}

func (s SkipAction) DefaultParam() []interface{} { return []interface{}{} }

func (s SkipAction) RandomParam() []interface{} { return []interface{}{} }

func (s SkipAction) Desc() string { return "# Skip action skips the current iteration" }

func (s SkipAction) Name() string { return "skip" }

func (s SkipAction) MaxParam() int { return 0 }

func (s SkipAction) MinParam() int { return 0 }

func (s SkipAction) ActionType() AType { return AnyAction }

type ExitAction struct{}

func (e ExitAction) DefaultParam() []interface{} { return []interface{}{} }

func (e ExitAction) RandomParam() []interface{} { return []interface{}{} }

func (e ExitAction) Desc() string { return "# Exit action exits the current function" }

func (e ExitAction) Name() string { return "exit" }

func (e ExitAction) MaxParam() int { return 0 }

func (e ExitAction) MinParam() int { return 0 }

func (e ExitAction) ActionType() AType { return AnyAction }
