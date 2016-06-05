package controller

import (
	"fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	// "github.com/go-macaron/session"
	// "macaron/controller"
	"strconv"
	"time"
	"zypc_submit/models"
)

var user = new(models.User)

func Loginhandler(ctx *macaron.Context) {
	ctx.Data["WebSiteIcon"] = websiteicon
	ctx.HTML(200, "login_demo")

}

func LoginJudgehandler(ctx *macaron.Context) (err error) {
	usernameinfo := ctx.Req.FormValue("userinfo")
	if usernameinfo[1] >= '0' && usernameinfo[1] <= '9' {
		user.UserId, err = strconv.ParseInt(usernameinfo, 10, 64)
		if err != nil {
			return err
		}
	} else {
		user.UserName = usernameinfo
	}
	user.Password = ctx.Req.FormValue("passwd")
	fmt.Println(user)

	if ok, _ := models.JudgeUser(user); ok {
		ctx.Redirect("/", 301)
	} else {
		ErrorInfo = "用户名错误或者密码错误！"
		ctx.Redirect("/errorinfo", 301)

	}
	return nil
}

func Registerhandler(ctx *macaron.Context) (err error) {

	userid := ctx.Req.FormValue("userid")
	user.UserId, err = strconv.ParseInt(userid, 10, 64)

	if err != nil {
		return err
	}

	user.UserName = ctx.Req.FormValue("username")
	user.Password = ctx.Req.FormValue("passwd")
	user.Email = ctx.Req.FormValue("email")
	user.Telnumber = ctx.Req.FormValue("telnumber")
	user.Time = time.Now()
	fmt.Println(user)

	if !NoUser(user.UserId) {
		err = models.AddUser(user)
	} else {
		ErrorInfo = "用户已存在，请查看用户信息填写是否正确无误，如果无误，请联系管理员。"
		ctx.Redirect("/errorinfo", 301)
	}

	if err != nil {
		return err
	}
	ctx.Redirect("/login", 301)
	return nil
}

func NoUser(userid int64) bool {
	has, _ := models.CheckUser(userid)

	return has
}
