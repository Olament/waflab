package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "waflab",
	Short: "WAFLab is a framework for testing Web Application Firewall",
	Long:  "",
}

var testcaseCount int64
var randomSeed int64
var omitIncompatible bool
var regexPerf bool

var compatibleList []string

func init() {
	rootCommand.PersistentFlags().Int64VarP(&randomSeed, "seed", "s", 41, "Define the seed used for generated each testcase. This can be used for debugging")
	rootCommand.Flags().BoolVar(&omitIncompatible, "omit-incompatible", true, "omit potentially incompatible configuration")
	rootCommand.Flags().BoolVar(&regexPerf, "performance", false, "if AutoGen should attempts to generate longest possible string from Regular Expression")

	generateCommand.Flags().Int64VarP(&testcaseCount, "count", "c", 1, "number of testcase generated for each variable")
	compatibleList = []string{"920", "921", "930", "931", "932", "933", "934", "941", "942", "943"}
}

func Execute() {
	rootCommand.AddCommand(generateCommand, testCommand)

	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
