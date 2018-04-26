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
	START  = 1
	Middle = 2
	END    = 3
)

type Node struct {
	name   string
	status int
	p      *Node
}

type Words struct {
	root map[string]Node
}

func (this *Words) Init(path string) {

	dir_list, _ := ioutil.ReadDir(path)

	for _, v := range dir_list {

		if v.Name()[0:1] == "." {
			continue
		}

		logs.Info(path+v.Name(), "reading")
		f, err := os.Open(path + v.Name())
		defer f.Close()

		if err != nil {
			logs.Warn(path, ":", err)
		}
		content, err := ioutil.ReadAll(f)
		logs.Info(path+v.Name(), "end")

		contentStr := strings.Trim(string(content), "*")

		a := strings.Split(contentStr, "\r\n")

		for i := 0; i < len(a); i++ {
			this.Add(strings.Trim(a[i], "*"))
		}
	}
}

func (this *Words) Add(word string) {
	wRune := []rune(word)
	fmt.Println(word)

	var node *Node
	for i := 0; i < len(wRune); i++ {

		tmp := wRune[i : i+1]

		node = new(Node)
		node.name = string(tmp)
		node.status = 0
		fmt.Println(string(tmp))
	}

}

func (this *Words) Del() {

}

func (this *Words) Update() {
	fmt.Println(123)
}
