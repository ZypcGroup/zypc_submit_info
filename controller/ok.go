package controller

import (
	// "fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	// "github.com/go-macaron/session"
	// "macaron/controller"
	// "strconv"
	// "time"
	// "zypc_submit/models"
)

func Okhandler(ctx *macaron.Context) {
	ctx.HTML(200, "ok")
}
