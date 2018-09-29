package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
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

func GenResultSet(packageInfo map[string]packageInfo) resultSet {
	rs := resultSet{
		[]Category{Category{}},
		[]eNode{eNode{}},
		[]eLink{eLink{}},
	}
	jsonByte, _ := json.Marshal(rs)
	fmt.Println(string(jsonByte))
	return rs
}
