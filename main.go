package main

import (
	"fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	// "github.com/go-macaron/session"
	// "macaron/controller"
	"zypc_submit/models"
	// "macaron/modules/initConf"
)

func init() {
	err := models.RegisterDB()
	if err != nil {
		fmt.Println("Error : ", err)
	}
}

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

	info := new(models.Infomation)

	info.Info1 = ctx.Req.FormValue("Info1")
	info.Info2 = ctx.Req.FormValue("Info2")
	info.Info3 = ctx.Req.FormValue("Info3")
	info.Info4 = ctx.Req.FormValue("Info4")
	info.Info5 = ctx.Req.FormValue("Info5")
	info.Info6 = ctx.Req.FormValue("info6")
	info.Info7 = ctx.Req.FormValue("info7")
	info.Info8 = ctx.Req.FormValue("info8")
	info.Info9 = ctx.Req.FormValue("info9")
	info.Info10 = ctx.Req.FormValue("info10")
	info.Info11 = ctx.Req.FormValue("info11")
	info.Info12 = ctx.Req.FormValue("info12")
	info.Info13 = ctx.Req.FormValue("info13")
	info.Info14 = ctx.Req.FormValue("info14")
	info.Info15 = ctx.Req.FormValue("info15")
	info.Info16 = ctx.Req.FormValue("info16")
	info.Info17 = ctx.Req.FormValue("info17")
	info.Info18 = ctx.Req.FormValue("info18")
	info.Info19 = ctx.Req.FormValue("info19")
	info.Info20 = ctx.Req.FormValue("info20")

	models.AddInfomation(info)
	ctx.Redirect("/", 301)
}
