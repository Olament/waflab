package yaml

import "github.com/waflab/waflab/test"

func DefaultYAML() *test.Testfile {
	v := &test.Testfile{
		Meta:  &test.Meta{
			Author:      "msra",
			Enabled:     false,
			Name:        "",
		},
		Tests: []*test.Test{
			{
				TestTitle: "",
				Desc: "",
				Stages: []*test.StageWrapper{
					{
						Stage: DefaultStage(),
					},
				},
			},
		},
	}
	return v
}

func DefaultStage() *test.Stage {
	v := &test.Stage{
		Input: &test.Input{
			SaveCookie: false,
			DestAddr:   "127.0.0.1",
			Method:     "GET",
			Port:       80,
			Protocol:   "http",
			Uri:        "/",
			Version:    "HTTP/1.1",
			Headers:    map[string]string{},
		},
		Output: &test.Output{
			Status:        []int{ 200 },
			HtmlContains:  "",
			LogContains:   "",
			NoLogContains: "",
		},
	}
	return v
}
