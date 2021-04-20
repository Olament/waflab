// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package autogen

import (
	"regexp"
	"strings"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
	"github.com/waflab/waflab/autogen/operator"
	"github.com/waflab/waflab/autogen/utils"
)

const (
	exclusionMaxRetry = 10
)

func processIndexExclusion(variables []*parser.Variable) ([]*parser.Variable, error) {
	// group exclusion of each variable together
	exclusions := make(map[int][]*parser.Variable)
	for _, v := range variables {
		if v.Exclusion {
			value := exclusions[v.Tk]
			exclusions[v.Tk] = append(value, v)
		}
	}

	newVariables := make([]*parser.Variable, 0)
	for _, v := range variables {
		if !v.Exclusion {
			isValidIndex := false
			// keep retry until we reach maximum retry time or obtain a valid index
			for i := 0; i < exclusionMaxRetry && !isValidIndex; i++ {
				// attempt to generate index
				var index string
				if v.Index == "" {
					// no index specified, give a random index
					index = utils.RandomString(10)
				} else {
					if isRegexIndex(v.Index) {
						// specified a regex index
						var err error
						index, err = operator.GenerateStringFromRegex(v.Index, false, operator.NoFlag)
						if err != nil {
							return nil, err
						}
					} else {
						// specified a literal index
						index = v.Index
					}
				}
				// check if index is valid
				isValidIndex = true
				for _, exclusion := range exclusions[v.Tk] {
					exclusionString := strings.Trim(exclusion.Index, `/`)
					matched, _ := regexp.MatchString("^"+exclusionString+"$", index)
					isValidIndex = isValidIndex && !matched
				}
				// construct new variables if we have a valid index
				if isValidIndex {
					newVariables = append(newVariables, &parser.Variable{
						Tk:    v.Tk,
						Index: index,
					})
				}
			}
		}
	}

	return newVariables, nil
}

// isRegexIndex checks if the index is a regex index
// Ex: for variable ARGS:abc, abc is not a regex index
//     however, ARGS:/ab*c/, /ab*c/ is a regex index
func isRegexIndex(index string) bool {
	return strings.HasPrefix(index, "/") && strings.HasSuffix(index, "/")
}
