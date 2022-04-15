package libs

import (
	"os"
	"os/exec"
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/core/logs"

	"github.com/fsnotify/fsnotify"
)

var (
	senWord Words
)

func Init() {

	logs.SetLogger(logs.AdapterFile, `{"filename":"ck.log"}`)

	path, _ := os.Getwd()
	path = path + "/" + beego.AppConfig.String("word_dir") + "/"
	if _, err := os.Stat(path); os.IsNotExist(err) {

		logs.Info("%s:file does not exist", path)
		os.MkdirAll(path, os.ModePerm)
		logs.Info("%s:make file", path)
	}

	//Initialize the word library
	senWord.Init(path)

	logs.Info("find start")
	s := senWord.Find("测试文本:在十九世纪，学系号,嘘唏不已欧洲,指纹套")
	logs.Info("r:", s)
	logs.Info("find end")

	logs.Info("monitor start")
	StartMonitor(path)
	logs.Info("monitor end")
}

func FindWord(text string) []string {
	return senWord.Find(text)
}

func StartMonitor(path string) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logs.Warn(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case e := <-watcher.Events:

				// if e.Op&fsnotify.Create == fsnotify.Create {
				// 	logs.Warn("Create:", e.Name)
				// }

				// if e.Op&fsnotify.Remove == fsnotify.Remove {
				// 	logs.Warn("Remove:", e.Name)
				// }

				// if e.Op&fsnotify.Rename == fsnotify.Rename {
				// 	logs.Warn("文件Rename:", e.Name)
				// }

				// if e.Op&fsnotify.Write == fsnotify.Write {
				// 	logs.Warn("文件Write:", e.Name)
				// }

				if e.Op&fsnotify.Chmod == fsnotify.Chmod {
					senWord.ReInit(path)
				}

			case err = <-watcher.Errors:
				if err != nil {
					logs.Warn("错误:", err)
				}

			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		logs.Warn("Failed to watch directory: ", err)
	}
	beego.Run()
	<-done
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
