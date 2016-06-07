package controller

import (
	"gopkg.in/macaron.v1"
	"zypc_submit/models"
)

func Showhandler(ctx *macaron.Context) {
	ctx.Data["Info"] = models.ShowInfomation()
	ctx.HTML(200, "show")
}
