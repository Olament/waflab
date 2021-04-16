// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package transformer

import (
	"fmt"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
	"github.com/waflab/waflab/autogen/utils"
)

type transformReverser func(variable string) string

var reverserFactory = map[int]transformReverser{
	parser.TkTransBase64Decode:       reverseBase64Decode,
	parser.TkTransCssDecode:          reverseCSSDecode,
	parser.TkTransCompressWhitespace: reverseCompressWhiteSpace,
	parser.TkTransHexDecode:          reverseHexDecode,
	parser.TkTransHtmlEntityDecode:   reverseHTMLEntityDecode,
	parser.TkTransJsDecode:           reverseJSDecode,
	parser.TkTransLength:             reverseLength,
	parser.TkTransNone:               reverseNone,
	parser.TkTransNormalisePath:      reverseNormalizePath,
	parser.TkTransNormalizePath:      reverseNormalizePath,
	parser.TkTransNormalisePathWin:   reverseNormalizePathWin,
	parser.TkTransNormalizePathWin:   reverseNormalizePathWin,
	parser.TkTransLowercase:          reverseLowercase,
	parser.TkTransRemoveComments:     reverseRemoveComments,
	parser.TkTransRemoveCommentsChar: reverseRemoveCommentsChar,
	parser.TkTransReplaceComments:    reverseReplaceComments,
	parser.TkTransRemoveNulls:        reverseRemoveNulls,
	parser.TkTransReplaceNulls:       reverseReplaceNulls,
	parser.TkTransTrim:               reverseTrim,
	parser.TkTransTrimLeft:           reverseTrimLeft,
	parser.TkTransTrimRight:          reverseTrimRight,
	parser.TkTransUrlDecode:          reverseUrlDecode,
	parser.TkTransUrlDecodeUni:       reverseUrlDecode,
	parser.TkTransUtf8toUnicode:      reverseUtf8ToUnicode,
}

func ReverseTransform(transformers []*parser.Trans, variable string) string {
	for i := len(transformers) - 1; i >= 0; i-- {
		if f, ok := reverserFactory[transformers[i].Tk]; ok {
			variable = f(variable)
		} else {
			fmt.Printf("\033[31m%s\033[0m\n", &utils.ErrNotSupported{
				Type: "transformation",
				Name: parser.TransformationNameMap[transformers[i].Tk],
			})
		}
	}
	return variable
}
