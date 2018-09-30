package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"io/ioutil"
	"os"
)

type resultSet struct {
	Categories []Category `json:"categories"`
	Nodes      []eNode    `json:"nodes"`
	Links      []eLink    `json:"links"`
}

type Category struct {
	Name string `json:"name"`
}

type eNode struct {
	ID         int    `json:"id"`
	Category   int    `json:"category"`
	Name       string `json:"name"`
	Label      string `json:"label"`
	SymbolSize int    `json:"symbolSize"`
	Ignore     bool   `json:"ignore"`
	Flag       bool   `json:"flag"`
}

type eLink struct {
	Source int `json:"source"`
	Target int `json:"target"`
}

// TODO genResultSet 生成 echarts 需要的结构体格式
// TODO 过滤非内部的依赖关系
func genResultSet(packageInfoMap map[string]packageInfo) resultSet {

	rs := resultSet{
		[]Category{Category{}},
		[]eNode{eNode{}},
		[]eLink{eLink{}},
	}
	jsonByte, _ := json.Marshal(rs)
	fmt.Println(string(jsonByte))
	return rs
}

// loadTemplate 加载页面模板
func loadTemplate() ([2]string, error) {
	f, err := os.Open(TEMPLATE_PATH)
	if err != nil {
		return [2]string{}, err
	}
	tByte, err := ioutil.ReadAll(f)
	if err != nil {
		return [2]string{}, err
	}

	return [2]string{string(tByte[:280]), string(tByte[635:])}, nil
}

// TODO generateView 生成展示页面
func generateView() {

}
