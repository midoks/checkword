package libs

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"os/exec"
	"strings"
)

var (
	senWord Words
)

func Init() {

	logs.SetLogger("file", `{"filename":"logs/ck.log"}`)

	path, _ := os.Getwd()
	path = path + "/words/"
	if _, err := os.Stat(path); os.IsNotExist(err) {

		logs.Info("%s:file does not exist", path)
		os.MkdirAll(path, os.ModePerm)
		logs.Info("%s:make file", path)
	}

	//Initialize the word library
	senWord.Init(path)

	logs.Info("find start")
	s := senWord.Find("在十九世纪，学系号,嘘唏不已欧洲")

	fmt.Println(s)
	logs.Info("find end")
}

func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
