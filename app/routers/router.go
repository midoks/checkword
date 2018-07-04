package routers

import (
	"github.com/astaxie/beego"
	"github.com/midoks/checkword/app/controllers"
	"net/http"
)

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("page not found"))
}

func init() {
	beego.ErrorHandler("404", page_not_found)

	ns := beego.NewNamespace("/v1", beego.NSAutoRouter(&controllers.IndexController{}))
	beego.AddNamespace(ns)
}
