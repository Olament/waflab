// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package autogen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/rule"
	y "gopkg.in/yaml.v2"
)

// GenerateTestFromDirectory read ModSecurity rule from dirPath and
// then write the generated test as yaml file into the ouput directory.
// It will create a folder with the name of config to hold all generated .yaml test file.
// Notice that it only read rules file ends with .conf
func GenerateTestFromDirectory(dirPath, output string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error %v at %q\n", err, path)
			return nil
		}

		// only process file with conf suffix
		if info.Mode().IsRegular() && filepath.Ext(path) == ".conf" {
			ruleStrings, err := ReadRuleStringFromConf(path)
			if err != nil {
				fmt.Printf("error %v when read %q\n", err, path)
				return nil
			}

			// generate testfiles from rules
			operator.WorkingDirectory = dirPath
			tests := GenerateTests(ruleStrings, 1)

			// write generated tests into files
			testDir := filepath.Join(filepath.Dir(output), info.Name())
			os.MkdirAll(testDir, os.ModePerm) // make a directory with the name of config
			for _, test := range tests {
				out, err := y.Marshal(test)
				if err != nil {
					fmt.Printf("error %v when marshal %s\n", err, test.Meta.Name)
					continue
				}
				err = ioutil.WriteFile(filepath.Join(testDir, test.Meta.Name), out, os.ModePerm)
				if err != nil {
					fmt.Printf("error %v when write %s\n", err, test.Meta.Name)
				}
			}

		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

// readRuleStringFromConf read rule string from config and
// remove any additional comments
func ReadRuleStringFromConf(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return rule.FilterSecRule(string(content)), nil
}
