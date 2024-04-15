package generate

import (
	"github.com/spf13/cobra"
	"github.com/tsinghua-cel/strategy-gen/command"
	"github.com/tsinghua-cel/strategy-gen/command/generate/config"
	"github.com/tsinghua-cel/strategy-gen/command/generate/export"
	"github.com/tsinghua-cel/strategy-gen/types"
)

func GetCommand() *cobra.Command {
	genCmd := &cobra.Command{
		Use:     "generate",
		Short:   "The default command to generate strategy.json with random strategies",
		PreRunE: runPreRun,
		Run:     runCommand,
	}

	registerSubcommands(genCmd)
	setFlags(genCmd)
	return genCmd
}

func setFlags(cmd *cobra.Command) {
	defaultConfig := config.DefaultConfig()
	cmd.Flags().StringVar(
		&params.configPath,
		configFlag,
		"",
		"the path to the CLI config. Supports .json and .yml",
	)

	cmd.Flags().IntVar(
		&params.rawConfig.ValidatorCount,
		validatorCountFlag,
		defaultConfig.ValidatorCount,
		"the number of validators",
	)

	cmd.Flags().IntVar(
		&params.rawConfig.StartSlot,
		startSlotFlag,
		defaultConfig.StartSlot,
		"the start slot to generate",
	)

	cmd.Flags().IntVar(
		&params.rawConfig.EndSlot,
		endSlotFlag,
		defaultConfig.EndSlot,
		"the end slot to generate",
	)

	cmd.Flags().StringVar(
		&params.rawConfig.EnableAttPoints,
		enableAttFlag,
		defaultConfig.EnableAttPoints,
		"the enabled attestation points, split with comma",
	)

	// EnableBlockPoints
	cmd.Flags().StringVar(
		&params.rawConfig.EnableBlockPoints,
		enableBlockFlag,
		defaultConfig.EnableBlockPoints,
		"the enabled block points, split with comma",
	)

	// EnableAttActions
	cmd.Flags().StringVar(
		&params.rawConfig.EnableAttActions,
		enableAttActionFlag,
		defaultConfig.EnableAttActions,
		"the enabled attestation actions, split with comma",
	)

	// EnableBlockActions
	cmd.Flags().StringVar(
		&params.rawConfig.EnableBlockActions,
		enableBlockActionFlag,
		defaultConfig.EnableBlockActions,
		"the enabled block actions, split with comma",
	)
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// server export
		export.GetCommand(),
	)
}

func runPreRun(cmd *cobra.Command, _ []string) error {
	// Check if the config file has been specified
	// Config file settings will override JSON-RPC and GRPC address values
	if isConfigFileSpecified(cmd) {
		if err := params.initConfigFromFile(); err != nil {
			return err
		}
	}

	return nil
}

func isConfigFileSpecified(cmd *cobra.Command) bool {
	return cmd.Flags().Changed(configFlag)
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	conf := params.rawConfig

	if err := runGenerate(conf); err != nil {
		outputter.SetError(err)
		outputter.WriteOutput()

		return
	}
}

func runGenerate(conf *config.Config) error {
	outputname := "strategy.json"
	strategy := types.Strategy{}
	valEndIndex := conf.ValidatorCount - 1
	validators := types.GetValidatorStrategy(0, valEndIndex, conf.StartSlot, conf.EndSlot)
	//slotStrategys := make([]types.SlotStrategy, 0)
	//for slot := conf.StartSlot; slot <= conf.EndSlot; slot++ {
	//	attIdx := rand.Intn(len(attestActions))
	//	bkIdx1 := rand.Intn(len(blockActions) - 1)
	//	bkIdx2 := bkIdx1 + rand.Intn(len(blockActions)-bkIdx1-1) + 1
	//	strategy := getSlotStrategy(fmt.Sprintf("%d", slot), attestActions[attIdx], blockActions[bkIdx1], blockActions[bkIdx2])
	//	slotStrategys = append(slotStrategys, strategy)
	//}
	//strategy.Slots = slotStrategys

	strategy.Validators = validators
	strategy.ToFile(outputname)

	return nil
}
