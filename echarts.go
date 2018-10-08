package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"io/ioutil"
	"os"
	"strings"
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

// genResultSet 生成 echarts 需要的结构体格式
func genResultSet(packageInfoMap map[string]packageInfo) resultSet {

	rs := resultSet{
		[]Category{},
		[]eNode{},
		[]eLink{},
	}

	idx := 0
	nodeMap := map[string]int{}

	// 添加主节点
	for _, pi := range packageInfoMap {
		if pi.ImportPath == LOCAL_PACKAGE_PATH {
			continue
		}
		nodeName := strings.TrimPrefix(pi.ImportPath, LOCAL_PACKAGE_PATH+SPILT)
		rs.Nodes = append(rs.Nodes, eNode{
			idx, idx, nodeName, nodeName, 40, false, true,
		})
		nodeMap[nodeName] = idx
		idx++
	}

	// 添加节点关系
	for _, pi := range packageInfoMap {
		for _, dep := range pi.Imports {

			// 过滤根
			if pi.ImportPath == LOCAL_PACKAGE_PATH {
				continue
			} // 过滤 go 内部包依赖关系
			if !strings.HasPrefix(dep, LOCAL_PACKAGE_PATH) {
				continue
			} // 过滤内部 vender 的依赖关系
			if strings.HasPrefix(dep, LOCAL_PACKAGE_PATH+"/vendor") {
				continue
			}

			depName := strings.TrimPrefix(dep, LOCAL_PACKAGE_PATH+SPILT)
			nodeName := strings.TrimPrefix(pi.ImportPath, LOCAL_PACKAGE_PATH+SPILT)
			nodeID := nodeMap[nodeName]

			if depID, ok := nodeMap[depName]; ok {
				rs.Links = append(rs.Links, eLink{nodeID, depID})
			} else {
				rs.Nodes = append(rs.Nodes, eNode{
					idx, idx, depName, depName, 30, false, true,
				})
				rs.Links = append(rs.Links, eLink{nodeID, idx})
				nodeMap[depName] = idx
				idx++
			}
		}
	}

	return rs
}

// loadTemplate 加载页面模板
func loadTemplate() ([2][]byte, error) {
	f, err := os.Open(TEMPLATE_PATH)
	if err != nil {
		return [2][]byte{}, err
	}
	tByte, err := ioutil.ReadAll(f)
	if err != nil {
		return [2][]byte{}, err
	}
	return [2][]byte{tByte[:280], tByte[632:]}, nil
}

// TODO 展示页面未来写死在代码里，减少加载文件可能出现的异常与耗时
// generateView 生成展示页面
func generateView(rs resultSet) string {
	jsonByte, _ := json.Marshal(rs)

	viewFilePath := WORKSPACE + "/godepView.html"
	f, err := os.OpenFile(viewFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer f.Close()
	f.Write(TEMPLATE[0])
	f.Write(jsonByte)
	f.Write(TEMPLATE[1])

	return viewFilePath
}
