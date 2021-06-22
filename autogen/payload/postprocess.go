// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package payload

import (
	"errors"
	"strings"

	"github.com/waflab/waflab/test"
)

var ErrReject = errors.New("autogen/postprocess: handler rejects the payload")

type filter func(payload *test.Input) error

var filters = []filter{
	crlfFilter,
	cookieFilter,
	xmlFilter,
}

// postprocess testing payload from AddVariable
func postprocess(payload *test.Input) error {
	for _, f := range filters {
		if err := f(payload); err != nil {
			return err
		}
	}
	return nil
}

// substitute Change Line and Line Feed character in Cookie
// as space character
// this filter always return nil
func crlfFilter(payload *test.Input) error {
	if _, okay := payload.Headers["Cookie"]; okay {
		replacer := strings.NewReplacer(
			"\n", " ",
			"\r", " ",
		)
		payload.Headers["Cookie"] = replacer.Replace(payload.Headers["Cookie"].(string))
	}
	return nil
}

// return ErrReject if cookie is invalid
func cookieFilter(payload *test.Input) error {
	if value, okay := payload.Headers["Cookie"]; okay {
		value := value.(string)
		if index := strings.IndexAny(value, ";"); index > 0 && index < len(value)-1 {
			// ; appears within the cookie body
			return ErrReject
		}
		if freq := strings.Count(value, "="); freq > 1 {
			// = appears more than once
			return ErrReject
		}
	}
	return nil
}

// xmlFilter replace the � (replacement character), which generate from converting unsupported
// white space character to XML encoding, to space character.
func xmlFilter(payload *test.Input) error {
	if payload.Data != nil {
		for index := 0; index < len(payload.Data); index++ {
			payload.Data[index] = strings.ReplaceAll(payload.Data[index], "�", " ")
		}
	}
	return nil
}
