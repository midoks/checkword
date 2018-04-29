package libs

import (
	// "bytes"
	"fmt"
	"os"
	// "math"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"strings"
)

const (
	NODE_START = 1
	NODE_MID   = 2
	NODE_END   = 3
)

type Words struct {
	root map[string]interface{}
}

func (this *Words) Init(path string) {

	this.root = make(map[string]interface{})

	dir_list, _ := ioutil.ReadDir(path)

	logs.Info("begin create")
	for _, v := range dir_list {

		if v.Name()[0:1] == "." {
			continue
		}

		f, err := os.Open(path + v.Name())
		defer f.Close()

		if err != nil {
			logs.Warn(path, ":", err)
		}

		content, err := ioutil.ReadAll(f)

		contentStr := strings.Trim(string(content), "*")
		contentStr = strings.Replace(contentStr, "\r\n", "\n", -1)

		a := strings.Split(contentStr, "\n")

		// logs.Warn(a)
		for i := 0; i < len(a); i++ {
			this.Add(strings.Trim(a[i], "*"))
		}
	}
	logs.Info("end create")
}

func (this *Words) Add(word string) {
	wRune := []rune(word)
	wlen := len(wRune)

	nowNode := this.root
	for i, thisWord := range wRune {

		tmpStr := string(thisWord)
		thisNode, ok := nowNode[tmpStr]

		if ok {

			if wlen-1 == i {
				thisNode.(map[string]interface{})["status"] = NODE_END
			}

			nowNode = thisNode.(map[string]interface{})
		} else {

			newNode := make(map[string]interface{})
			nowNode[tmpStr] = newNode

			if wlen-1 == i {
				newNode["status"] = NODE_END
			} else if i == 0 {
				newNode["status"] = NODE_START
			} else {
				newNode["status"] = NODE_MID
			}

			nowNode = newNode
		}
	}
	// logs.Info("root:", this.root)
}

func (this *Words) Find(str string) []string {
	r := []rune(str)
	llen := len(r)

	var ss string = ""
	var result []string

	nowNode := this.root

	for i := 0; i < llen; i++ {

		tmpWord := string(r[i])
		if thisNode, ok := nowNode[tmpWord]; ok {

			ss = ss + tmpWord
			nowNode = thisNode.(map[string]interface{})

			if i == llen-1 {
				if thisNodeStatus, ok := nowNode["status"]; ok {
					if thisNodeStatus == NODE_END {
						result = append(result, ss)
					}
				}
			}
		} else {

			if len([]rune(ss)) >= 1 {
				i--
				if thisNodeStatus, ok := nowNode["status"]; ok {
					if thisNodeStatus == NODE_END {
						result = append(result, ss)
					}
				}
			}

			ss = ""
			nowNode = this.root
		}
	}
	return result
}

func (this *Words) Del(str string) {

}

func (this *Words) Update() {
	fmt.Println(123)
}
