package routers

import (
	"github.com/astaxie/beego"
	"github.com/midoks/checkword/app/controllers"
)

func init() {

	ns := beego.NewNamespace("/v1", beego.NSAutoRouter(&controllers.IndexController{}))
	beego.AddNamespace(ns)
}
