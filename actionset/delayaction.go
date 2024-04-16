package actionset

import "math/rand"

var (
	defaultDelayWithSecondAction Action = delayWithSecondAction{
		config: ActionConfig{
			Name:                "delayWithSecond",
			Random:              false,
			ParamCount:          1,
			DefaultParamValue:   4,
			MinRandomParamValue: 0,
			MaxRandomValue:      15,
		},
	}

	defaultDelayToNextSlotAction Action = delayToNextSlotAction{
		config: ActionConfig{
			Name:                "delayToNextSlot",
			Random:              false,
			ParamCount:          0,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}

	defaultDelayToAfterNextSlotAction Action = delayToAfterNextSlotAction{
		config: ActionConfig{
			Name:                "delayToAfterNextSlot",
			Random:              false,
			ParamCount:          1,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}

	defaultDelayToNextNEpochStartAction Action = delayToNextNEpochStartAction{
		config: ActionConfig{
			Name:                "delayToNextNEpochStart",
			Random:              false,
			ParamCount:          1,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}

	defaultDelayToNextNEpochEndAction Action = delayToNextNEpochEndAction{
		config: ActionConfig{
			Name:                "delayToNextNEpochEnd",
			Random:              false,
			ParamCount:          1,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}

	defaultDelayToNextNEpochHalfAction Action = delayToNextNEpochHalfAction{
		config: ActionConfig{
			Name:                "delayToNextNEpochHalf",
			Random:              false,
			ParamCount:          1,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}

	defaultDelayToEpochEndAction Action = delayToEpochEndAction{
		config: ActionConfig{
			Name:                "delayToEpochEnd",
			Random:              false,
			ParamCount:          0,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}

	defaultDelayHalfEpochAction Action = delayHalfEpochAction{
		config: ActionConfig{
			Name:                "delayHalfEpoch",
			Random:              false,
			ParamCount:          0,
			DefaultParamValue:   0,
			MinRandomParamValue: 0,
			MaxRandomValue:      0,
		},
	}
)

type delayWithSecondAction struct {
	config ActionConfig
}

func (d delayWithSecondAction) DefaultParam() []interface{} { return []interface{}{4} }

func (d delayWithSecondAction) RandomParam() []interface{} {
	r := rand.Intn(15)
	return []interface{}{r}
}

func (d delayWithSecondAction) Desc() string { return "# Delay with seconds" }

func (d delayWithSecondAction) Name() string { return "delayWithSecond" }

func (d delayWithSecondAction) MaxParam() int { return 1 }

func (d delayWithSecondAction) MinParam() int { return 1 }

func (d delayWithSecondAction) GetConfig() ActionConfig { return d.config }

func (d delayWithSecondAction) WithConfig(config ActionConfig) Action {
	d.config = config
	return d
}

func (d delayWithSecondAction) ActionType() AType { return AnyAction }

type delayToNextSlotAction struct {
	config ActionConfig
}

func (d delayToNextSlotAction) DefaultParam() []interface{} { return []interface{}{} }

func (d delayToNextSlotAction) RandomParam() []interface{} { return []interface{}{} }

func (d delayToNextSlotAction) Desc() string { return "# Delay to next slot" }

func (d delayToNextSlotAction) Name() string { return "delayToNextSlot" }

func (d delayToNextSlotAction) MaxParam() int { return 0 }

func (d delayToNextSlotAction) MinParam() int { return 0 }

func (d delayToNextSlotAction) GetConfig() ActionConfig { return d.config }

func (d delayToNextSlotAction) WithConfig(config ActionConfig) Action {
	d.config = config
	return d
}

func (d delayToNextSlotAction) ActionType() AType { return AnyAction }

type delayToAfterNextSlotAction struct {
	config ActionConfig
}

func (d delayToAfterNextSlotAction) DefaultParam() []interface{} { return []interface{}{0} }

func (d delayToAfterNextSlotAction) RandomParam() []interface{} { return []interface{}{0} }

func (d delayToAfterNextSlotAction) Desc() string { return "# Delay to after next slot" }

func (d delayToAfterNextSlotAction) Name() string { return "delayToAfterNextSlot" }

func (d delayToAfterNextSlotAction) MaxParam() int { return 1 }

func (d delayToAfterNextSlotAction) MinParam() int { return 1 }

func (d delayToAfterNextSlotAction) GetConfig() ActionConfig { return d.config }

func (d delayToAfterNextSlotAction) WithConfig(config ActionConfig) Action {
	d.config = config
	return d
}

func (d delayToAfterNextSlotAction) ActionType() AType { return AnyAction }

type delayToNextNEpochStartAction struct {
	config ActionConfig
}

func (d delayToNextNEpochStartAction) DefaultParam() []interface{} { return []interface{}{0} }

func (d delayToNextNEpochStartAction) RandomParam() []interface{} { return []interface{}{0} }

func (d delayToNextNEpochStartAction) Desc() string { return "# Delay to next n epoch start" }

func (d delayToNextNEpochStartAction) Name() string { return "delayToNextNEpochStart" }

func (d delayToNextNEpochStartAction) MaxParam() int { return 1 }

func (d delayToNextNEpochStartAction) MinParam() int { return 1 }

func (d delayToNextNEpochStartAction) GetConfig() ActionConfig { return d.config }

func (d delayToNextNEpochStartAction) WithConfig(config ActionConfig) Action {
	d.config = config
	return d
}

func (d delayToNextNEpochStartAction) ActionType() AType { return AnyAction }

type delayToNextNEpochEndAction struct {
	config ActionConfig
}

func (d delayToNextNEpochEndAction) DefaultParam() []interface{} { return []interface{}{0} }

func (d delayToNextNEpochEndAction) RandomParam() []interface{} { return []interface{}{0} }

func (d delayToNextNEpochEndAction) Desc() string { return "# Delay to next n epoch end" }

func (d delayToNextNEpochEndAction) Name() string { return "delayToNextNEpochEnd" }

func (d delayToNextNEpochEndAction) MaxParam() int { return 1 }

func (d delayToNextNEpochEndAction) MinParam() int { return 1 }

func (d delayToNextNEpochEndAction) GetConfig() ActionConfig { return d.config }

func (d delayToNextNEpochEndAction) WithConfig(config ActionConfig) Action {
	d.config = config
	return d
}

func (d delayToNextNEpochEndAction) ActionType() AType { return AnyAction }

type delayToNextNEpochHalfAction struct {
	config ActionConfig
}

func (d delayToNextNEpochHalfAction) DefaultParam() []interface{} { return []interface{}{0} }

func (d delayToNextNEpochHalfAction) RandomParam() []interface{} { return []interface{}{0} }

func (d delayToNextNEpochHalfAction) Desc() string { return "# Delay to next n epoch half" }

func (d delayToNextNEpochHalfAction) Name() string { return "delayToNextNEpochHalf" }

func (d delayToNextNEpochHalfAction) MaxParam() int { return 1 }

func (d delayToNextNEpochHalfAction) MinParam() int { return 1 }

func (d delayToNextNEpochHalfAction) GetConfig() ActionConfig { return d.config }

func (d delayToNextNEpochHalfAction) WithConfig(config ActionConfig) Action {
	d.config = config
	return d
}

func (d delayToNextNEpochHalfAction) ActionType() AType { return AnyAction }

type delayToEpochEndAction struct {
	config ActionConfig
}

func (d delayToEpochEndAction) DefaultParam() []interface{} { return []interface{}{} }

func (d delayToEpochEndAction) RandomParam() []interface{} { return []interface{}{} }

func (d delayToEpochEndAction) Desc() string { return "# Delay to epoch end" }

func (d delayToEpochEndAction) Name() string { return "delayToEpochEnd" }

func (d delayToEpochEndAction) MaxParam() int { return 0 }

func (d delayToEpochEndAction) MinParam() int { return 0 }

func (d delayToEpochEndAction) GetConfig() ActionConfig { return d.config }

func (d delayToEpochEndAction) WithConfig(config ActionConfig) Action {
	d.config = config
	return d
}

func (d delayToEpochEndAction) ActionType() AType { return AnyAction }

type delayHalfEpochAction struct {
	config ActionConfig
}

func (d delayHalfEpochAction) DefaultParam() []interface{} { return []interface{}{} }

func (d delayHalfEpochAction) RandomParam() []interface{} { return []interface{}{} }

func (d delayHalfEpochAction) Desc() string { return "# Delay half epoch" }

func (d delayHalfEpochAction) Name() string { return "delayHalfEpoch" }

func (d delayHalfEpochAction) MaxParam() int { return 0 }

func (d delayHalfEpochAction) MinParam() int { return 0 }

func (d delayHalfEpochAction) GetConfig() ActionConfig { return d.config }

func (d delayHalfEpochAction) WithConfig(config ActionConfig) Action {
	d.config = config
	return d
}

func (d delayHalfEpochAction) ActionType() AType { return AnyAction }
