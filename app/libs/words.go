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

		// logs.Info("reading")
		f, err := os.Open(path + v.Name())
		defer f.Close()

		if err != nil {
			logs.Warn(path, ":", err)
		}

		content, err := ioutil.ReadAll(f)
		// logs.Info("end")

		contentStr := strings.Trim(string(content), "*")

		a := strings.Split(contentStr, "\n")

		for i := 0; i < len(a); i++ {
			// fmt.Println(strings.Trim(a[i], "*"), "-----------|")
			this.Add(strings.Trim(a[i], "*"))
		}
	}
	logs.Info("end create")
}

func (this *Words) Add(word string) {
	wRune := []rune(word)

	nowNode := this.root
	wlen := len(wRune)

	for i, thisWord := range wRune {

		tmpStr := string(thisWord)

		thisNode, ok := nowNode[tmpStr]

		if ok {

			nowNode = thisNode.(map[string]interface{})
		} else {

			newNode := make(map[string]interface{})
			newNode["end"] = false
			nowNode[tmpStr] = newNode

			nowNode = newNode
		}

		if i == (wlen - 1) {
			nowNode["end"] = true
		}
	}
}

func (this *Words) Find(str string) []string {
	r := []rune(str)
	llen := len(r)

	var ss string = ""
	var result []string

	nowNode := this.root

	fmt.Println(nowNode)

	for i := 0; i < llen; i++ {
		tmpWord := string(r[i])
		thisNode, ok := nowNode[tmpWord]

		fmt.Println(thisNode)

		if ok {
			tmpNode := thisNode.(map[string]interface{})
			if tmpNode["end"].(bool) {

				if i+1 <= llen {
					tmp2 := string(r[i+1])
					thisNode2, ok2 := tmpNode[tmp2]
					fmt.Println("ee start:", tmp2, thisNode2)

					if ok2 {

						tmpNode2 := thisNode2.(map[string]interface{})
						if tmpNode2["end"].(bool) {
							continue
						}
					}
					fmt.Println("ee end:", tmp2)

				}

				ss = ss + tmpWord
				result = append(result, ss)
				ss = ""
				nowNode = this.root
			} else {
				ss = ss + tmpWord
				nowNode = tmpNode
			}
		} else {

			if len([]rune(ss)) >= 1 {
				i--
			}
			ss = ""
			nowNode = this.root
		}
	}
	return result
}

func (this *Words) Del() {

}

func (this *Words) Update() {
	fmt.Println(123)
}
