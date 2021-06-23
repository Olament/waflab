package command

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/object"
)

var generateCommand = &cobra.Command{
	Use:   "generate [WAF rule files directory] [output directory]",
	Short: "Generate ftw-compatible YAML test file from WAF rules",
	Args:  cobra.MinimumNArgs(0),
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
	var ruleFileDirectory string
	var outputDirectory string

	if len(args) == 0 { // use default if no parameter supplied
		object.InitRepo()
		ruleFileDirectory = path.Join("repos", "coreruleset", "rules")
		fmt.Println(ruleFileDirectory)
		outputDirectory = path.Join("output")
	}

	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, "Need specify rule file directory and output directory")
		return
	}

	operator.WorkingDirectory = ruleFileDirectory
	testfiles, err := generateTestfile(ruleFileDirectory, int(testcaseCount))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = writeTestfile(outputDirectory, testfiles)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
