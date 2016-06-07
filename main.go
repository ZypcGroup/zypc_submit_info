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

	Sess, _ = session.NewManager("memory", session.Options{Provider: "memory"})
	fmt.Println(Sess)
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
	m.Get("/test2", Test2handler)
	m.Run(port)
}

func Testhandler(ctx *macaron.Context, f *session.Flash) {
	// sess.Set("session", "axiu session")
	// sessid, _ := Sess.RegenerateId(ctx)
	// fmt.Println(sessid.ID())

	sess, _ := Sess.Start(ctx)

	ct := sess.Get("Count")
	if ct == nil {
		sess.Set("Count", 1)
	} else {
		sess.Set("Count", (ct.(int) + 1))
	}

	fmt.Println(ct)
	ctx.Data["Count"] = ct
	f.Success("yes!!!")
	f.Error("opps...")
	f.Info("aha?!")
	f.Warning("Just be careful.")
	ctx.HTML(200, "test")
}

func Test2handler(ctx *macaron.Context, f *session.Flash) {
	sess, _ := Sess.Start(ctx)
	// sess.Set("Count", 1)
	ct := sess.Get("Count")
	fmt.Println(ct)
	ctx.Data["Count"] = ct
	ctx.HTML(200, "test")

}
