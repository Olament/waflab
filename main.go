// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/waflab/waflab/object"
	"github.com/waflab/waflab/routers"
)

func main() {
	object.InitRepo()
	object.InitMaster()
	object.InitOrmManager()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//beego.DelStaticPath("/static")
	beego.SetStaticPath("/static", "web/build/static")
	beego.InsertFilter("/", beego.BeforeRouter, routers.TransparentStatic) // must has this for default page
	beego.InsertFilter("/*", beego.BeforeRouter, routers.TransparentStatic)

	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 * 24 * 365

	beego.Run()
}
