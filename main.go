package main

import (
	"fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	// "github.com/go-macaron/session"
	// "macaron/controller"
	// "macaron/models"
	// "macaron/modules/initConf"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", mainhandler)
	m.Post("/submit", submithandler)
	m.Run()

}

func mainhandler(ctx *macaron.Context) {
	ctx.Data["IsLogin"] = false
	ctx.Data["Info1"] = "学号"
	ctx.Data["Info2"] = "班级"
	ctx.Data["Info3"] = "姓名"
	ctx.Data["Info4"] = "博客地址"
	ctx.Data["Info5"] = "备注"

	ctx.HTML(200, "home")
}

func submithandler(ctx *macaron.Context) {

	info1 := ctx.Req.FormValue("Info1")
	info2 := ctx.Req.FormValue("Info2")
	info3 := ctx.Req.FormValue("Info3")
	info4 := ctx.Req.FormValue("Info4")
	info5 := ctx.Req.FormValue("Info5")

	fmt.Println(info1, info2, info3, info4, info5)
	ctx.Redirect("/", 301)
}
