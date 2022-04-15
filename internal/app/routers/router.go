package routers

import (
	"net/http"

	"github.com/astaxie/beego"

	"github.com/midoks/checkword/internal/app/controllers"
)

func page_not_found(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("page not found"))
}

func init() {
	beego.ErrorHandler("404", page_not_found)

	ns := beego.NewNamespace("/v1", beego.NSAutoRouter(&controllers.IndexController{}))
	beego.AddNamespace(ns)
}
