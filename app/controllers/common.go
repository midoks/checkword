package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
	"time"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type CommonController struct {
	beego.Controller
	RequestTime time.Time
}

func (this *CommonController) Prepare() {
	this.startRecTimes()
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
	out["time"] = this.CalcLoadTimes()

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
		this.StopRun()
	}
}

func (this *CommonController) startRecTimes() time.Time {
	this.RequestTime = time.Now()
	return this.RequestTime
}

func (this *CommonController) CalcLoadTimes() string {
	return fmt.Sprintf("%dms", (time.Now().Sub(this.RequestTime).Nanoseconds() / 1e6))
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
