package runtime

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tsinghua-cel/strategy-gen/library"
	"github.com/tsinghua-cel/strategy-gen/types"
)

func GetCommand() *cobra.Command {
	runtimeCmd := &cobra.Command{
		Use:   "runtime",
		Short: "Run a library strategy generator.",
		Run:   runCommand,
	}
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	setFlags(runtimeCmd)
	return runtimeCmd
}

func setFlags(cmd *cobra.Command) {

	cmd.Flags().IntVar(
		&params.maxValidatorIndex,
		maxValidatorIndexFlag,
		-1,
		"the max hack validator index",
	)

	cmd.Flags().StringVar(
		&params.attacker,
		attackerFlag,
		"127.0.0.1:12001",
		"the attacker service to update",
	)

	cmd.Flags().StringVar(
		&params.strategy,
		strategyFlag,
		"",
		"runtime strategy to used",
	)

	cmd.Flags().BoolVar(
		&params.listLibrary,
		listFlag,
		false,
		"list all library strategies",
	)
}

func runCommand(cmd *cobra.Command, _ []string) {
	library.Init()

	if params.listLibrary {
		listLibrary()
		return
	}
	if params.strategy == "" {
		log.Fatal("strategy is required")
	}

	strategy, ok := library.GetStrategy(params.strategy)
	if !ok {
		log.Fatalf("strategy %s not found", params.strategy)
	}

	strategy.Run(types.LibraryParams{
		Attacker:          params.attacker,
		MaxValidatorIndex: params.maxValidatorIndex,
	})

}

func listLibrary() {
	strategies := library.GetAllStrategies()
	for name, strategy := range strategies {
		fmt.Printf("%s:\n %s\n", name, strategy.Description())
	}
}
