package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	"github.com/go-macaron/session"
	"zypc_submit/controller"
	"zypc_submit/models"
	// "macaron/modules/initConf"
)

const (
	Port int = 8080
)

var (
	port int = Port
)

var conf *goconfig.ConfigFile

var Sess *session.Manager

func init() {

	err := models.RegisterDB()
	if err != nil {
		fmt.Println("Error : ", err)
	}

	conf, err = goconfig.LoadConfigFile("conf/app.conf")
	if err != nil {
		fmt.Println("Load Config File Error! \t", err)
	}

	if ok := conf.MustInt("Server", "ListenPort"); ok != 0 {
		port = ok
	}
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Get("/", controller.Homehandler)
	m.Post("/submit", controller.Submithandler)
	m.Get("/login", controller.Loginhandler)
	m.Post("/login", controller.LoginJudgehandler)
	m.Post("/register", controller.Registerhandler)
	m.Get("/errorinfo", controller.ErrorInfohandler)
	// m.Get("/test", controller.Testhandler)
	m.Get("/test", Testhandler)
	m.Run(port)
}

func Testhandler(ctx *macaron.Context, f *session.Flash) {
	// sess.Set("session", "axiu session")

	f.Success("yes!!!")
	f.Error("opps...")
	f.Info("aha?!")
	f.Warning("Just be careful.")
	ctx.HTML(200, "test")
}
