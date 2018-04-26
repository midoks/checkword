package controllers

import (
	_ "github.com/astaxie/beego"
)

type MainController struct {
	CommonController
}

// logs.GetLogger("ORM").Println("this is a message of orm")
// logs.Debug("my book is bought in the year of ", 2016)
// logs.Info("this %s cat is %v years old", "yellow", 3)
// logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
// logs.Error(1024, "is a very", "good game")
// logs.Critical("oh,crash")
func (this *MainController) Get() {

	out := make(map[string]interface{})
	out["Website"] = "beego.me"
	out["Email"] = "astaxie@gmail.com"

	test := "i你好"
	nameRune := []rune(test)

	out["p"] = test[0:3]
	out["p1"] = string(nameRune[0:1])

	this.retJson(out)
}
