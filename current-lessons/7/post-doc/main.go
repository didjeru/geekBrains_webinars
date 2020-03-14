package main

import (
	_ "not-for-work/GeekBrainsWebinars/current-lessons/7/post-doc/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run("localhost:8080")
}