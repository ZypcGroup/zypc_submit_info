package controller

import (
	"fmt"
	// "github.com/Unknwon/goconfig"
	// "github.com/go-macaron/gzip"
	"gopkg.in/macaron.v1"
	// "log"
	"github.com/go-macaron/session"
	// "macaron/controller"
	// "encoding/base64"
	"strconv"
	"time"
	"zypc_submit/models"
)

var Sess *session.Manager

var user = new(models.User)

func init() {
	Sess, _ = session.NewManager("memory", session.Options{Provider: "memory"})
	// fmt.Println(Sess)

}

func Loginhandler(ctx *macaron.Context) {
	ctx.Data["WebSiteIcon"] = websiteicon
	ctx.HTML(200, "login_demo")

}

func LoginJudgehandler(ctx *macaron.Context) (err error) {
	user = nil
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

		sess, _ := Sess.Start(ctx)

		createtime := sess.Get("CreateTime")
		if createtime == nil {
			sess.Set("CreateTime", time.Now().Unix())
			sess.Set("UserID", user.UserId)
			// sess.Set("Countnum", 1)
			ctx.Redirect("/", 301)
		} else if (createtime.(int64) + 360) > time.Now().Unix() {
			// fmt.Println(createtime, "\n-------------\n")

			ctx.Data["IsLogin"] = true
			ctx.Redirect("/", 301)
		} else {
			ctx.Redirect("/login", 301)
		}

	} else {
		ErrorInfo = "用户名错误或者密码错误！"
		ctx.Redirect("/errorinfo", 301)

	}
	return nil
}

func Registerhandler(ctx *macaron.Context) (err error) {
	user = nil
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
	// fmt.Println(user)

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
	user = nil
	has, _ := models.CheckUser(userid)

	return has
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

	// sessid := sess.Get("MacaronSession")

	// fmt.Println(ct, sessid)

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
	// fmt.Println(ct)
	ctx.Data["Count"] = ct
	ctx.HTML(200, "test")

}
