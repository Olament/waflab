package command

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/docker"
	"github.com/waflab/waflab/parse"
	"github.com/waflab/waflab/test"
	"github.com/waflab/waflab/util"
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
var jsonPath string

func init() {
	testCommand.Flags().BoolVar(&noHost, "no-host", false, "stop appending host header of target address to testcase")
	testCommand.Flags().StringVarP(&confDirectory, "config", "g", "", "generate testcases from WAF rules")
	testCommand.Flags().StringVarP(&yamlDirectory, "yaml", "y", "", "read testcases from yaml files")
	testCommand.Flags().StringVar(&format, "format", "%NAME | %STATUS | %HIT | %EXPECTED_STAT | %HT_MATCH | %STAT_MATCH", "indicate the format of result")
	testCommand.Flags().StringVar(&filter, "filter", ".*", "specify a regular expression to filter out hitrules")
	testCommand.Flags().StringVarP(&jsonPath, "json", "j", "repos/wafrules-drs-2.0.json", "read enabled rules in the specified json file")
}

func appendTestcases(testcases []string, title string, yamlFile string, enabledRules map[string]bool) []string {
	if enabledRules != nil {
		if _, ok := enabledRules[title]; ok {
			testcases = append(testcases, yamlFile)
			fmt.Printf("%s | included in enabled rules | added\n", title)
		} else {
			fmt.Printf("%s | not included in enabled rules | abandoned\n", title)
		}
	} else {
		testcases = append(testcases, yamlFile)
	}

	return testcases
}

func testing(cmd *cobra.Command, args []string) {
	if confDirectory == "" && yamlDirectory == "" {
		if _, err := os.Stat("output"); !os.IsNotExist(err) {
			yamlDirectory = "output"
		} else {
			confDirectory = path.Join("repos", "coreruleset", "rules")
		}
	}

	var yamlTestcases []string

	// get enabled rules map from wafrules.json or wafrules-drs-2.0.json
	var enabledRuleSet map[string]bool
	if util.FileExist(jsonPath) {
		enabledRuleSet = parse.GetEnabledRules(jsonPath)
	}

	titleStatusMap := map[string][]int{} // title to corresponding status code
	if confDirectory != "" {             // generate testcase from config
		operator.WorkingDirectory = confDirectory
		testcases, err := generateTestfile(confDirectory, int(testcaseCount))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		for _, test := range testcases {
			t := test.Tests[0]
			strs := strings.Split(t.TestTitle, "-")
			testTitle := strs[0]
			for _, t := range test.Tests {
				for index, stage := range t.Stages {
					if !noHost {
						if stage.Stage.Input.Headers == nil {
							stage.Stage.Input.Headers = make(map[string]interface{})
						}
						stage.Stage.Input.Headers["Host"] = args[0]
					}
					if index == 0 {
						titleStatusMap[t.TestTitle] = stage.Stage.Output.Status
					}
				}
			}
			out, err := yaml.Marshal(&test)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			yamlTestcases = appendTestcases(yamlTestcases, testTitle, string(out), enabledRuleSet)
		}
	} else { // read from yaml directory
		err := filepath.Walk(yamlDirectory, func(path string, info fs.FileInfo, err error) error {
			if info.Mode().IsRegular() && filepath.Ext(path) == ".yaml" {
				out, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}
				t := test.Testfile{}
				yaml.Unmarshal(out, &t)
				for _, t := range t.Tests {
					for index, stage := range t.Stages {
						if index == 0 {
							titleStatusMap[t.TestTitle] = stage.Stage.Output.Status
						}
					}
				}
				re := regexp.MustCompile(`(\d+)\.yaml`)
				strs := re.FindStringSubmatch(path)
				testTitle := strs[1]
				yamlTestcases = appendTestcases(yamlTestcases, testTitle, string(out), enabledRuleSet)
			}
			return nil
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}

	master := docker.MakeMaster(5)
	re := regexp.MustCompile(filter)
	numsOfTestcases := len(yamlTestcases)
	finishedTestcases := 0

	if err := ui.Init(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer ui.Close()

	list := widgets.NewList()
	list.SelectedRowStyle = ui.NewStyle(ui.ColorYellow)

	progress := widgets.NewGauge()
	progress.BarColor = ui.ColorCyan

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(ui.NewRow(9.0/10, list), ui.NewRow(1.0/10, progress))

	startTime := time.Now()

	uiEvents := ui.PollEvents()

	for _, y := range yamlTestcases {
		results, err := master.InsertTask(args[0], y)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		for _, res := range results {
			var expectedStatus string = ""
			var hitMatch string = ""
			var statusMatch string = ""
			if status, okay := titleStatusMap[res.Title]; okay {
				expectedStatus = fmt.Sprint(status)
			}
			if strings.Contains(expectedStatus, res.Status) && res.Status != "" {
				statusMatch = "YES"
			} else {
				statusMatch = "NO"
			}
			if strings.Contains(res.HitRule, strings.Split(res.Title, "-")[0]) {
				hitMatch = "YES"
			} else {
				hitMatch = "NO"
			}
			replacer := strings.NewReplacer(
				"%NAME", res.Title,
				"%STATUS", res.Status,
				"%HIT", strings.Join(re.FindAllString(res.HitRule, -1), " "),
				"%EXPECTED_STAT", expectedStatus,
				"%HT_MATCH", hitMatch,
				"%STAT_MATCH", statusMatch,
			)
			list.Rows = append(list.Rows, replacer.Replace(format))
			list.ScrollBottom()
		}
		curTime := time.Now()
		finishedTestcases++
		progress.Percent = int((float64(finishedTestcases) / float64(numsOfTestcases)) * 100)
		progress.Label = fmt.Sprintf("%d/%d, %.2f rps", finishedTestcases, numsOfTestcases, float64(finishedTestcases)/float64(curTime.Sub(startTime).Seconds()))
		ui.Render(grid)

		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		default:
			// do nothing
		}
	}

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Down>":
			list.ScrollDown()
		case "<Up>":
			list.ScrollUp()
		}

		ui.Render(list)
	}
}
