package routers

import (
	"github.com/midoks/checkword/app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
