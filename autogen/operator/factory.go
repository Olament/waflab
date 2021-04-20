// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package operator

import (
	"github.com/hsluoyz/modsecurity-go/seclang/parser"
	"github.com/waflab/waflab/autogen/utils"
)

// flag
const (
	NoFlag = iota
	NoUTF8
)

type operationReverser func(argument string, not bool) (string, error)
type operationReverserWithFlag func(argument string, not bool, flag int) (string, error)

var reverserFactory = map[int]operationReverser{
	// string matching operator
	parser.TkOpBeginsWith: reverseBeginsWith,
	parser.TkOpContains:   reverseContains,
	parser.TkOpEndsWith:   reverseEndsWith,
	parser.TkOpPm:         reversePm,
	parser.TkOpPmFromFile: reversePmFromFile,
	parser.TkOpStrEq:      reverseStrEq,
	parser.TkOpWithin:     reverseWithin,
	// numerical operator
	parser.TkOpEq: reverseEq,
	parser.TkOpGe: reverseGe,
	parser.TkOpGt: reverseGt,
	parser.TkOpLe: reverseLe,
	parser.TkOpLt: reverseLt,
	// validation operator
	parser.TkOpValidateByteRange:    reverseValidateByteRange,
	parser.TkOpValidateUtf8Encoding: reverseValidateUtf8Encoding,
	parser.TkOpValidateUrlEncoding:  reverseValidateURLEncoding,
	// miscellaneous operator
	parser.TkOpIpMatch:         reverseIPMatch,
	parser.TkOpIpMatchFromFile: reverseIPMatchFromFile,
	parser.TkOpDetectSqli:      reverseDetectSQLi,
	parser.TkOpDetectXss:       reverseDetectXSS,
}

var reverserFactoryWithFlag = map[int]operationReverserWithFlag{
	parser.TkOpRx: reverseRx,
}

// ReverseOperator generate a string by reversing the given ModSecurity Operator.
func ReverseOperator(operator *parser.Operator, flag int) (string, error) {
	if f, ok := reverserFactory[operator.Tk]; ok {
		return f(operator.Argument, operator.Not)
	}
	if f, ok := reverserFactoryWithFlag[operator.Tk]; ok {
		return f(operator.Argument, operator.Not, flag)
	}
	return "", &utils.ErrNotSupported{
		Type: "operator",
		Name: parser.OperatorNameMap[operator.Tk],
	}
}
