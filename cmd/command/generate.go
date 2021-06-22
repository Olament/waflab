package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/waflab/waflab/autogen/operator"
)

var generateCommand = &cobra.Command{
	Use:   "generate [WAF rule files directory] [output directory]",
	Short: "Generate ftw-compatible YAML test file from WAF rules",
	Args:  cobra.MinimumNArgs(1),
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
	operator.WorkingDirectory = args[0]
	testfiles, err := generateTestfile(args[0], int(testcaseCount))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	err = writeTestfile(args[1], testfiles)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
