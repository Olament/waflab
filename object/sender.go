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
				Status:   -2,
				Response: "No connection",
			}
			return res
		}
		defer resp.Body.Close()

		res := &Result{}
		res.Status = resp.StatusCode
		if util.IntListContains(expectedStatusList, res.Status) {
			res.Response = "ok"
		} else {
			res.Response = fmt.Sprintf("wrong: %v != %d", expectedStatusList, res.Status)
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

func sendRaw(method string, host string, uri string, query string, userAgent string, headers map[string]string) (*http.Response, error) {
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
			req.Header.Add(k, v)
		} else {
			req.Header.Add(k, v)
		}
	}

	return client.Do(req)
}
