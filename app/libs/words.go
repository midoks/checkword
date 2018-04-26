package libs

import (
	// "bytes"
	// "fmt"
	// "math"
	// "strings"
	"github.com/astaxie/beego/logs"
)

type Node struct {
	name   string
	status int
	p      *Node
}

type Words struct {
	node []Node
}

func (this *Words) Init(path string) {
	// os.
	logs.Info(path)
}

func (this *Words) Add() {
}

func (this *Words) Del() {

}

func (this *Words) Update() {

}
