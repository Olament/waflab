// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package controllers

import (
	"github.com/waflab/waflab/docker"
	"github.com/waflab/waflab/object"
)

func (c *ApiController) GetResult() {
	testsetId := c.Input().Get("testsetId")
	testcaseId := c.Input().Get("testcaseId")
	typ := c.Input().Get("type")

	c.Data["json"] = object.GetResult(testsetId, testcaseId, typ)
	c.ServeJSON()
}

func (c *ApiController) GetDockerHealth() {
	c.Data["json"] = &struct {
		Health bool `json:"health"`
	}{
		docker.GetMasterHealth(),
	}
	c.ServeJSON()
}
