// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package parse

import (
	"encoding/json"
	"strings"

	"github.com/waflab/waflab/util"
)

func GetEnabledRules(path string) map[string]bool {
	s := util.ReadStringFromPath(path)
	rulesFile := WafRulesFile{}
	err := json.Unmarshal([]byte(s), &rulesFile)
	if err != nil {
		panic(err)
	}

	enabledRuleSet := make(map[string]bool)

	for _, ruleSet := range rulesFile.RuleSets {
		if !strings.Contains(ruleSet.Name, "DefaultRuleSet_2.0") {
			continue
		}

		ruleGroups := ruleSet.Properties.RuleGroups
		for _, ruleGroup := range ruleGroups {
			for _, rule := range ruleGroup.Rules {
				if rule.DefaultState == "Enabled" {
					enabledRuleSet[rule.RuleId] = true
				}
			}
		}
	}

	return enabledRuleSet
}
