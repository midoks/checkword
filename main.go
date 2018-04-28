package main

import (
	"github.com/astaxie/beego"
	"github.com/midoks/checkword/app/libs"
	_ "github.com/midoks/checkword/app/routers"
)

func main() {
	libs.Init()
	beego.Run()
}
