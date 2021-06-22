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

func init() {
	rootCommand.PersistentFlags().Int64VarP(&randomSeed, "seed", "s", 41, "Define the seed used for generated each testcase. This can be used for debugging")
	generateCommand.Flags().Int64VarP(&testcaseCount, "count", "c", 1, "number of testcase generated for each variable")
}

func Execute() {
	rootCommand.AddCommand(generateCommand, testCommand)

	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
