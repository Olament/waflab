// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package payload

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/waflab/waflab/test"
)

func composeCookie(payload *test.Input, key string, value string) {
	composeHeader(payload, "Cookie", fmt.Sprintf("%s=%s", key, value))
}

func composeQueryString(payload *test.Input, key string, value string) {
	if value == "" {
		payload.Uri = fmt.Sprintf("/?%s", url.QueryEscape(key))
	} else {
		payload.Uri = fmt.Sprintf("/?%s=%s", url.QueryEscape(key), url.QueryEscape(value))
	}
}

func composeHeader(payload *test.Input, key string, value string) {
	payload.Headers[key] = value
}

func composeFile(payload *test.Input, name, filename, content string) {
	composeHeader(payload, "Content-Type", "multipart/form-data; boundary=X-BOUNDARY")
	composeHeader(payload, "Cache-Control", "no-cache")
	composeHeader(payload, "Host", "localhost")
	payload.Method = "POST"
	payload.Data = []string{
		"--X-BOUNDARY",
		fmt.Sprintf("Content-Disposition: form-data; name=\"%s\"; filename=\"%s\"",
			httpRequestEscape(name),
			httpRequestEscape(filename)),
		"Content-Type: text/plain",
		"",
		content,
		"",
		"--X-BOUNDARY--",
		"",
		"",
	}
	// calculate Content-Length
	length := 0
	for _, d := range payload.Data {
		length += len(d) + 2 // including /r/n
	}
	payload.Headers["Content-Length"] = strconv.Itoa(length - 4) // excluding trailing \r\n\r\n
}

// setURI set the Uri entry in test.Input
// value is value for Uri entry (slash "/" not included)
// setURI will attempt to format it as a valid Uri value
func setURI(payload *test.Input, value string) {
	if !isValidURL(value) {
		value = url.QueryEscape(value)
	}
	if _, err := url.ParseRequestURI(fmt.Sprintf("/%s", value)); err == nil {
		payload.Uri = fmt.Sprintf("/%s", value)
		return
	}
	payload.Uri = fmt.Sprintf("/?%s", value)
}
