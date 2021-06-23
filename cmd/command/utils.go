package command

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/waflab/waflab/autogen"
	"github.com/waflab/waflab/test"
	"gopkg.in/yaml.v2"
)

func generateTestfile(dir string, testcaseCount int) ([]*test.Testfile, error) {
	testcases := []*test.Testfile{}
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.Mode().IsRegular() && filepath.Ext(path) == ".conf" {
			if omitIncompatible && !checkCompatible(compatibleList, info.Name()) {
				return nil
			}
			ruleStrings, err := autogen.ReadRuleStringFromConf(path)
			if err != nil {
				return err
			}
			testcases = append(testcases, autogen.GenerateTests(ruleStrings, int(testcaseCount))...)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return testcases, nil
}

func writeTestfile(dir string, testfiles []*test.Testfile) error {
	os.MkdirAll(dir, os.ModePerm) // make a directory with the name of config
	for _, test := range testfiles {
		out, err := yaml.Marshal(test)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Join(dir, test.Meta.Name), out, os.ModePerm)
		if err != nil {
			fmt.Printf("error %v when write %s\n", err, test.Meta.Name)
		}
	}
	return nil
}

func checkCompatible(compatibleList []string, fileName string) bool {
	for _, c := range compatibleList {
		if strings.Contains(fileName, c) {
			return true
		}
	}
	return false
}
