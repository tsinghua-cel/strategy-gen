package config

import (
	"encoding/json"
	"fmt"
	"github.com/tsinghua-cel/strategy-gen/actionset"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

// Config defines the server configuration params
type Config struct {
	ValidatorCount     int                      `json:"validator_count" yaml:"validator_count"`
	StartSlot          int                      `json:"start_slot" yaml:"start_slot"`
	EndSlot            int                      `json:"end_slot" yaml:"end_slot"`
	EnableBlockPoints  string                   `json:"enable_block_points" yaml:"enable_block_points"`
	EnableBlockActions string                   `json:"enable_block_actions" yaml:"enable_block_actions"`
	EnableAttPoints    string                   `json:"enable_att_points" yaml:"enable_att_points"`
	EnableAttActions   string                   `json:"enable_att_actions" yaml:"enable_att_actions"`
	ActionsConfig      []actionset.ActionConfig `json:"actions_config" yaml:"actions_config"`
}

const (
	DefaultValidatorCount     = 21
	DefaultStartSlot          = 0
	DefaultEndSlot            = 10000
	DefaultEnableBlockPoints  = "BlockDelayForReceiveBlock,BlockBeforeBroadCastBlock"
	DefaultEnableAttPoints    = "AttestBeforeBroadCast"
	DefaultEnableBlockActions = "null,delayWithSecond,delayToAfterNextSlot,delayToNextNEpochStart,delayToNextNEpochHalf,delayToEpochEnd,return"
	DefaultEnableAttActions   = "null,delayWithSecond,delayToAfterNextSlot,return"
)

// DefaultConfig returns the default server configuration
func DefaultConfig() *Config {
	conf := &Config{
		ValidatorCount:     DefaultValidatorCount,
		StartSlot:          DefaultStartSlot,
		EndSlot:            DefaultEndSlot,
		EnableBlockPoints:  DefaultEnableBlockPoints,
		EnableBlockActions: DefaultEnableBlockActions,
		EnableAttPoints:    DefaultEnableAttPoints,
		EnableAttActions:   DefaultEnableAttActions,
		ActionsConfig:      make([]actionset.ActionConfig, 0),
	}
	allaction := actionset.GetAllActionSet()
	for _, action := range allaction {
		conf.ActionsConfig = append(conf.ActionsConfig, action.GetConfig())
	}
	return conf
}

// ReadConfigFile reads the config file from the specified path, builds a Config object
// and returns it.
//
// Supported file types: .json, .hcl, .yaml, .yml
func ReadConfigFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var unmarshalFunc func([]byte, interface{}) error

	switch {
	case strings.HasSuffix(path, ".json"):
		unmarshalFunc = json.Unmarshal
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		unmarshalFunc = yaml.Unmarshal
	default:
		return nil, fmt.Errorf("suffix of %s is neither hcl, json, yaml nor yml", path)
	}

	config := DefaultConfig()
	if err := unmarshalFunc(data, config); err != nil {
		return nil, err
	}

	return config, nil
}
