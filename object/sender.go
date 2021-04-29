// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package object

import (
	"fmt"
	"net/http"

	"github.com/waflab/waflab/util"
)

func getWafResult(testset *Testset, testcase *Testcase) *Result {
	tf := testcase.Data
	for _, test := range tf.Tests {
		stage := test.Stages[0].Stage
		input := stage.Input
		output := stage.Output

		method := input.Method
		host := testset.TargetUrl
		uri := input.Uri
		query := ""
		headers := input.Headers

		expectedStatusList := output.Status

		resp, err := sendRaw(method, host, uri, query, "", headers)
		if err != nil {
			//panic(err)
			res := &Result{
				Statuses: []int{-2},
				Response: "No connection",
			}
			return res
		}
		defer resp.Body.Close()

		res := &Result{}
		res.Statuses = []int{resp.StatusCode}
		if util.IntListContains(expectedStatusList, res.Statuses[0]) {
			res.Response = "ok"
		} else {
			res.Response = fmt.Sprintf("wrong: %v != %d", expectedStatusList, res.Statuses[0])
		}
		//if resp.StatusCode == http.StatusOK {
		//	bodyBytes, err := ioutil.ReadAll(resp.Body)
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	res.Response = string(bodyBytes)
		//}

		return res
	}

	return nil
}

func sendRaw(method string, host string, uri string, query string, userAgent string, headers map[string]interface{}) (*http.Response, error) {
	client := &http.Client{}
	//host = "http://127.0.0.1:8888"
	url := host + uri + query
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	if userAgent != "" {
		req.Header.Set("User-Agent", userAgent)
	}

	for k, v := range headers {
		// https://github.com/golang/go/issues/7682
		if k == "Host" {
			req.Host = ""
			req.Header.Add(k, v.(string))
		} else {
			req.Header.Add(k, v.(string))
		}
	}

	return client.Do(req)
}
