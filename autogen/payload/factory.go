// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package payload

import (
	"github.com/waflab/waflab/autogen/utils"
	"github.com/waflab/waflab/test"

	"github.com/hsluoyz/modsecurity-go/seclang/parser"
)

type payloadConverter func(value, index string, payload *test.Input) error

var converterFactory = map[int]payloadConverter{
	parser.TkVarArgs:                addArg,
	parser.TkVarArgsCombinedSize:    addArgCombinedSize,
	parser.TkVarArgsNames:           addArgNames,
	parser.TkVarArgsGet:             addArg,      // equivalent to ARGS
	parser.TkVarArgsGetNames:        addArgNames, // equivalent to ARGS_NAMES
	parser.TkVarExtendedJSON:        addExtendedJSON,
	parser.TkVarFiles:               addFiles,
	parser.TkVarFilesNames:          addFilesNames,
	parser.TkVarFilesCombinedSize:   addFilesCombinedSize,
	parser.TkVarQueryString:         addQueryString,
	parser.TkVarRequestBasename:     addFilesNames, // equivalent to FILES_NAMES
	parser.TkVarRequestBody:         addRequestBody,
	parser.TkVarRequestCookies:      addRequestCookies,
	parser.TkVarRequestCookiesNames: addRequestCookiesName,
	parser.TkVarRequestFilename:     addRequestFileName,
	parser.TkVarRequestHeaders:      addRequestHeaders,
	parser.TkVarRequestHeadersNames: addRequestHeadersNames,
	parser.TkVarRequestLine:         addRequestLine,
	parser.TkVarRequestMethod:       addRequestMethod,
	parser.TkVarRequestProtocol:     addRequestProtocol,
	parser.TkVarRequestUri:          addRequestURI,
	parser.TkVarRequestUriRaw:       addRequestURIRaw,
	parser.TkVarXML:                 addXML,
}

func AddVariable(v *parser.Variable, value string, payload *test.Input) error {
	if f, ok := converterFactory[v.Tk]; ok {
		err := f(value, v.Index, payload)
		return err
	}
	return &utils.ErrNotSupported{
		Type: "Variable",
		Name: parser.VariableNameMap[v.Tk],
	}
}
