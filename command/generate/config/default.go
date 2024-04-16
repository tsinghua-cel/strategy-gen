package config

import "github.com/tsinghua-cel/strategy-gen/actionset"

var (
	defaultNullActionConfig = actionset.ActionConfig{
		Name:                "null",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultReturnActionConfig = actionset.ActionConfig{
		Name:                "return",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultContinueActionConfig = actionset.ActionConfig{
		Name:                "continue",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultAbortActionConfig = actionset.ActionConfig{
		Name:                "abort",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultSkipActionConfig = actionset.ActionConfig{
		Name:                "skip",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultExitActionConfig = actionset.ActionConfig{
		Name:                "exit",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultDelayWithSecondActionConfig = actionset.ActionConfig{
		Name:                "delayWithSecond",
		Random:              false,
		ParamCount:          1,
		DefaultParamValue:   4,
		MinRandomParamValue: 0,
		MaxRandomValue:      15,
	}
	defaultDelayToNextSlotActionConfig = actionset.ActionConfig{
		Name:                "delayToNextSlot",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultDelayToAfterNextSlotActionConfig = actionset.ActionConfig{
		Name:                "delayToAfterNextSlot",
		Random:              false,
		ParamCount:          1,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultDelayToNextNEpochStartActionConfig = actionset.ActionConfig{
		Name:                "delayToNextNEpochStart",
		Random:              false,
		ParamCount:          1,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultDelayToNextNEpochEndActionConfig = actionset.ActionConfig{
		Name:                "delayToNextNEpochEnd",
		Random:              false,
		ParamCount:          1,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultStoreSignedAttestActionConfig = actionset.ActionConfig{
		Name:                "storeSignedAttest",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
	defaultRePackAttestationActionConfig = actionset.ActionConfig{
		Name:                "rePackAttestation",
		Random:              false,
		ParamCount:          0,
		DefaultParamValue:   0,
		MinRandomParamValue: 0,
		MaxRandomValue:      0,
	}
)
