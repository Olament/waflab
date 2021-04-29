// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package test

import (
	"errors"
	"strings"
)

type Testfile struct {
	Meta  *Meta   `yaml:"meta"`
	Tests []*Test `yaml:"tests"`
}

type Meta struct {
	Author      string `yaml:"author"`
	Enabled     bool   `yaml:"enabled"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Test struct {
	TestTitle string          `yaml:"test_title"`
	Desc      string          `yaml:"desc"`
	Stages    []*StageWrapper `yaml:"stages"`
}

type StageWrapper struct {
	Stage *Stage `yaml:"stage"`
}

type Stage struct {
	Input  *Input  `yaml:"input"`
	Output *Output `yaml:"output"`
}

// DataSlice is a customized type for 'data' entry in the Input YAML file
type DataSlice []string

// UnmarshalYAML is a customized parser that parse both YAML list or string into DataSlice
// since both of them are valid type for 'data' entry per ftw YAML file definition
func (data *DataSlice) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err == nil {
		*data = strings.Split(str, "\n")
		return nil
	}

	slice := make([]string, 0)
	if err := unmarshal(&slice); err == nil {
		*data = slice
		return nil
	}

	return errors.New("failed to unmarshal into DataSlice")
}

type Input struct {
	SaveCookie     bool                   `yaml:"save_cookie,omitempty"`
	StopMagic      bool                   `yaml:"stop_magic,omitempty"`
	DestAddr       string                 `yaml:"dest_addr"`
	Method         string                 `yaml:"method"`
	Port           int                    `yaml:"port"`
	Protocol       string                 `yaml:"protocol"`
	Uri            string                 `yaml:"uri"`
	Version        string                 `yaml:"version"`
	Headers        map[string]interface{} `yaml:"headers,omitempty"`
	Data           DataSlice              `yaml:"data,omitempty"`
	EncodedRequest string                 `yaml:"encoded_request,omitempty"`
	RawRequest     string                 `yaml:"raw_request,omitempty"`
}

type Output struct {
	Status        []int  `yaml:"status"`
	HtmlContains  string `yaml:"html_contains,omitempty"`
	LogContains   string `yaml:"log_contains,omitempty"`
	NoLogContains string `yaml:"no_log_contains,omitempty"`
}
