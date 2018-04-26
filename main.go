package main

import (
	_ "github.com/midoks/checkword/app/routers"
	"github.com/midoks/checkword/app/libs"
	"github.com/astaxie/beego"
)

func main() {
	libs.Init()
	beego.Run()
}

