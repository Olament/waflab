// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package utils

import (
	"fmt"
	"strings"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
)

var charset = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r",
	"s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_", "~", "@", "!", "#"}

// RandomStringFromSet generate a string by concating freq numbers of strings randomly draw from set.
// Notice that string draw from set may be repetitive.
func RandomStringFromSet(freq int, set []string) string {
	var builder strings.Builder
	for i := 0; i < freq; i++ {
		builder.WriteString(PickRandomString(set))
	}
	return builder.String()
}

// RandomString generate a random string with given length using character from charSet.
func RandomString(length int) string {
	return RandomStringFromSet(length, charset)
}

// PickRandomString picks a random string from the given string slice
func PickRandomString(set []string) string {
	return set[randomGenerator.rand.Intn(len(set))]
}

// RuleDump dumps debug string for rule
func RuleDump(r *parser.RuleDirective) string {
	var builder strings.Builder
	dumpAction(&builder, r.Actions)
	dumpVariable(&builder, r.Variable)
	dumpOperator(&builder, r.Operator)
	return builder.String()
}

func dumpVariable(b *strings.Builder, variables []*parser.Variable) {
	for _, variable := range variables {
		if variable.Count {
			b.WriteString("&")
		}
		if variable.Exclusion {
			b.WriteString("!")
		}
		b.WriteString(parser.VariableNameMap[variable.Tk])
		if variable.Index != "" {
			b.WriteString(":" + variable.Index)
		}
		b.WriteString("\n")
	}
}

func dumpAction(b *strings.Builder, action *parser.Actions) {
	b.WriteString(fmt.Sprintf("%d\n", action.Id))
	for _, ac := range action.Action {
		if ac.Tk == parser.TkActionAllow || ac.Tk == parser.TkActionBlock ||
			ac.Tk == parser.TkActionDeny || ac.Tk == parser.TkActionPass {
			b.WriteString(fmt.Sprintf("%s\n", parser.ActionNameMap[ac.Tk]))
		}
	}
	for _, t := range action.Trans {
		b.WriteString(fmt.Sprintf("t:%s ", parser.TransformationNameMap[t.Tk]))
	}
	b.WriteString("\n")
}

func dumpOperator(b *strings.Builder, operator *parser.Operator) {
	if operator.Not {
		b.WriteString("!")
	}
	b.WriteString(fmt.Sprintf("@%s %s\n", parser.OperatorNameMap[operator.Tk], operator.Argument))
}
