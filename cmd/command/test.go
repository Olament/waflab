package command

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/cobra"
	"github.com/waflab/waflab/docker"
	"gopkg.in/yaml.v2"
)

var testCommand = &cobra.Command{
	Use:   "test [target WAF address]",
	Short: "Testing WAF with testcase generated from WAF rules or existing testcases",
	Args:  cobra.MinimumNArgs(1),
	Run:   testing,
}

var noHost bool
var confDirectory string
var yamlDirectory string
var format string
var filter string

func init() {
	testCommand.Flags().BoolVar(&noHost, "no-host", false, "stop appending host header of target address to testcase")
	testCommand.Flags().StringVarP(&confDirectory, "config", "g", "", "generate testcases from WAF rules")
	testCommand.Flags().StringVarP(&yamlDirectory, "yaml", "y", "", "read testcases from yaml files")
	testCommand.Flags().StringVar(&format, "format", "%NAME | %STATUS | %HIT", "indicate the format of result")
	testCommand.Flags().StringVar(&filter, "filter", ".*", "specify a regular expression to filter out hitrules")
}

func testing(cmd *cobra.Command, args []string) {
	if confDirectory == "" && yamlDirectory == "" {
		fmt.Fprintln(os.Stderr, "You must specify a source of testcases using config or yaml flag!")
		return
	}

	var yamlTestcases []string
	if confDirectory != "" { // generate testcase from config
		testcases, err := generateTestfile(confDirectory, int(testcaseCount))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		for _, test := range testcases {
			if !noHost {
				for _, t := range test.Tests {
					for _, stage := range t.Stages {
						if stage.Stage.Input.Headers == nil {
							stage.Stage.Input.Headers = make(map[string]interface{})
						}
						stage.Stage.Input.Headers["Host"] = args[0]
					}
				}
			}
			out, err := yaml.Marshal(&test)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			yamlTestcases = append(yamlTestcases, string(out))
		}
	} else { // read from yaml directory
		err := filepath.Walk(yamlDirectory, func(path string, info fs.FileInfo, err error) error {
			if info.Mode().IsRegular() && filepath.Ext(path) == ".yaml" {
				out, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				yamlTestcases = append(yamlTestcases, string(out))
			}
			return nil
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	master := docker.MakeMaster(5)
	results := []docker.Response{}
	bar := pb.StartNew(len(yamlTestcases))
	for _, y := range yamlTestcases {
		res, err := master.InsertTask(args[0], y)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		results = append(results, res...)
		bar.Increment()
	}
	bar.Finish()

	re := regexp.MustCompile(filter)

	for _, res := range results {
		replacer := strings.NewReplacer(
			"%NAME", res.Title,
			"%STATUS", res.Status,
			"%HIT", strings.Join(re.FindAllString(res.HitRule, -1), " "),
		)
		fmt.Println(replacer.Replace(format))
	}
}
