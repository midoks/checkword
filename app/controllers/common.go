package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type CommonController struct {
	beego.Controller
}

// 是否POST提交
func (this *CommonController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

// 重定向
func (this *CommonController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

//获取用户IP地址
func (this *CommonController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

func (this *CommonController) retResult(code int, msg interface{}, data ...interface{}) {
	out := make(map[string]interface{})
	out["code"] = code
	out["msg"] = msg

	if len(data) > 0 {
		out["data"] = data
	}

	this.retJson(out)
}

func (this *CommonController) D(args ...string) {
	if beego.AppConfig.String("runmode") == "dev" {
		for i := 0; i < len(args); i++ {
			this.Ctx.WriteString(args[i])
		}
		//this.StopRun()
	}
}

// 输出json
func (this *CommonController) retJson(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *CommonController) retOk(msg interface{}, data ...interface{}) {
	this.retResult(MSG_OK, msg, data...)
}

func (this *CommonController) retFail(msg interface{}, data ...interface{}) {
	this.retResult(MSG_ERR, msg, data...)
}
