// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package autogen

import (
	"errors"
	"fmt"
	"log"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"

	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/autogen/payload"
	"github.com/waflab/waflab/autogen/transformer"
	"github.com/waflab/waflab/autogen/utils"
	"github.com/waflab/waflab/autogen/yaml"
	"github.com/waflab/waflab/rule"
	"github.com/waflab/waflab/test"
)

const (
	maxRetry = 10
)

func GenerateTests(ruleString string, maxRetry int) (YAMLs []*test.Testfile) {
	rules, err := rule.ParseRuleDataToList(ruleString)
	if err != nil {
		log.Printf("Err while parsing rule string %s\n", ruleString)
		return nil
	}

	for _, rule := range rules {
		if rule.Actions == nil || rule.Actions.Id == 0 {
			continue
		}
		if rule.Actions.Chain { // Chained rule
			log.Printf("Err chain rule %d not supported\n", rule.Actions.Id)
			continue
		}
		if t := processIndependentRule(rule, maxRetry); t != nil {
			YAMLs = append(YAMLs, t)
		}
	}
	return YAMLs
}

// processIndependentRule generate testcases from given independent rule
// the max number of test cases it can generate is (# of variable in rule) * targetCases
// the actual number of test generated will likely lower than the max number since there may be
// duplicated test case
func processIndependentRule(rule *parser.RuleDirective, targetCases int) *test.Testfile {
	res := yaml.DefaultYAML()

	// set meta information
	res.Meta.Author = "Microsoft"
	res.Meta.Name = fmt.Sprintf("dev-%d.yaml", rule.Actions.Id)
	res.Meta.Description = "This YAML file is automatically generated by WAFLab AutoGen"

	// process variable index exclusion
	newVariables, err := processIndexExclusion(rule.Variable)
	if err != nil {
		log.Printf("Rule %d: skip, %v", rule.Actions.Id, err)
		return nil
	}
	rule.Variable = newVariables

	/* rule generation with duplication check */
	isDuplicate := make(map[string]bool)
	for i := 0; i < targetCases; i++ {
		valid := false
		for j := 0; !valid && j < maxRetry; j++ {
			tests, signature, err := generateTestCaseSet(rule)
			if _, ok := err.(*utils.ErrNotSupported); ok { // rule not supported
				return nil
			}
			if errors.Is(err, payload.ErrReject) { // generation fail, retry
				continue
			}
			if err != nil { // unhandled error
				log.Printf("autogen/strategy: %v\n", err)
				return nil
			}
			if _, okay := isDuplicate[signature]; okay { // repeated generation
				continue
			}
			res.Tests = append(res.Tests, tests...)
			valid = true
			isDuplicate[signature] = true
		}
	}

	for index := 0; index < len(res.Tests); index++ {
		res.Tests[index].TestTitle = fmt.Sprintf("%d-%d", rule.Actions.Id, index)
	}

	return res
}

// generateTestCaseSet generates "a set" of testcase. For example, a rule with 3 variables, this method
// will generate three testcase, where each testcase corresponds to a variable. signature is the unique
// value for the set of testcase and will be used to filter out repeated testcases.
func generateTestCaseSet(rule *parser.RuleDirective) (tests []*test.Test, signature string, err error) {
	// prepare the testcase struct
	for i := 0; i < len(rule.Variable); i++ {
		tests = append(tests, &test.Test{
			Stages: []*test.StageWrapper{
				{
					Stage: yaml.DefaultStage(),
				},
			},
		})
	}
	// get status code
	statusCode := getStatusCode(rule)

	/* reversing operator */
	reversed := ""
	if rule.Actions.Id == 944200 {
		reversed, err = operator.ReverseOperator(rule.Operator, operator.NoUTF8)
	} else {
		reversed, err = operator.ReverseOperator(rule.Operator, operator.NoFlag)
	}
	if err != nil {
		return nil, "", err
	}

	/* reversing transformation */
	reversed = transformer.ReverseTransform(rule.Actions.Trans, reversed)

	/* add testcase for each variable */
	for index, variable := range rule.Variable {
		err = payload.AddVariable(variable, reversed, tests[index].Stages[0].Stage.Input)
		if err != nil {
			return nil, "", err
		}
		tests[index].Stages[0].Stage.Output.Status = statusCode
		tests[index].Desc = parser.VariableNameMap[variable.Tk]
	}

	return tests, reversed, nil
}

func getStatusCode(rule *parser.RuleDirective) (statusCode []int) {
	for _, action := range rule.Actions.Action {
		switch action.Tk {
		case parser.TkActionAllow, parser.TkActionPass:
			statusCode = []int{200, 404}
		case parser.TkActionDeny, parser.TkActionBlock:
			statusCode = []int{403}
		default:
		}
	}
	return statusCode
}
