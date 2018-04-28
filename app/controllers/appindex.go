package controllers

import (
	"encoding/json"
	// "fmt"
	_ "github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"github.com/midoks/checkword/app/libs"
)

type IndexController struct {
	CommonController
}

// logs.GetLogger("ORM").Println("this is a message of orm")
// logs.Debug("my book is bought in the year of ", 2016)
// logs.Info("this %s cat is %v years old", "yellow", 3)
// logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
// logs.Error(1024, "is a very", "good game")
// logs.Critical("oh,crash")

func (this *IndexController) Index() {
	out := make(map[string]interface{})
	out["web"] = "http://midoks.cacheche.com"
	out["mail"] = "midoks@163.com"
	this.retJson(out)
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {

		tmp := a[i]
		isFind := false
		for j := i + 1; j < a_len; j++ {
			if tmp == a[j] {
				isFind = true
				break
			}
		}

		if !isFind {
			ret = append(ret, a[i])
		}
	}
	return
}

func (this *IndexController) Check() {

	if this.isPost() {

		req := make(map[string]interface{})
		json.Unmarshal(this.Ctx.Input.RequestBody, &req)

		if req["text"] != nil {

			r := libs.FindWord(req["text"].(string))

			c := len(r)
			// logs.Info(req["text"], r)
			if c == 0 {
				this.retResult(0, "没有敏感词!")
			} else {
				this.retResult(1, RemoveDuplicatesAndEmpty(r))
			}
		}

		this.retResult(-1, "缺少text字段")
	}

	this.retResult(-1, "请求错误!")
}
