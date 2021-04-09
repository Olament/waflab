// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package routers

import (
	"github.com/astaxie/beego"

	"github.com/waflab/waflab/controllers"
)

func init() {
	initAPI()
}

func initAPI() {
	ns :=
		beego.NewNamespace("/api",
			beego.NSInclude(
				&controllers.ApiController{},
			),
		)
	beego.AddNamespace(ns)

	beego.Router("/api/list-rulesets", &controllers.ApiController{}, "GET:ListRulesets")
	beego.Router("/api/list-rulefiles", &controllers.ApiController{}, "GET:ListRulefiles")
	beego.Router("/api/list-rules", &controllers.ApiController{}, "GET:ListRules")

	beego.Router("/api/get-testsets", &controllers.ApiController{}, "GET:GetTestsets")
	beego.Router("/api/get-testset", &controllers.ApiController{}, "GET:GetTestset")
	beego.Router("/api/update-testset", &controllers.ApiController{}, "POST:UpdateTestset")
	beego.Router("/api/add-testset", &controllers.ApiController{}, "POST:AddTestset")
	beego.Router("/api/delete-testset", &controllers.ApiController{}, "POST:DeleteTestset")

	beego.Router("/api/get-testcases", &controllers.ApiController{}, "GET:GetTestcases")
	beego.Router("/api/get-filtered-testcases", &controllers.ApiController{}, "GET:GetFilteredTestcases")
	beego.Router("/api/get-testcase", &controllers.ApiController{}, "GET:GetTestcase")
	beego.Router("/api/update-testcase", &controllers.ApiController{}, "POST:UpdateTestcase")
	beego.Router("/api/add-testcase", &controllers.ApiController{}, "POST:AddTestcase")
	beego.Router("/api/delete-testcase", &controllers.ApiController{}, "POST:DeleteTestcase")

	beego.Router("/api/get-result", &controllers.ApiController{}, "GET:GetResult")
	beego.Router("/api/get-docker-health", &controllers.ApiController{}, "GET:GetDockerHealth")
}
