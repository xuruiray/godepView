package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"io/ioutil"
	"os"
	"strings"
)

type cResultSet struct {
	Nodes []data `json:"nodes"`
	Links []data `json:"edges"`
}

type data struct {
	Data interface{} `json:"data"`
}

type cNode struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type cLink struct {
	Source int `json:"source"`
	Target int `json:"target"`
}

// genEResultSet 生成 echarts 需要的结构体格式
func genCResultSet(packageInfoMap map[string]packageInfo) cResultSet {

	rs := cResultSet{
		[]data{},
		[]data{},
	}

	idx := 0
	nodeMap := map[string]int{}

	// 添加主节点
	for _, pi := range packageInfoMap {
		if pi.ImportPath == LOCAL_PACKAGE_PATH {
			continue
		}
		nodeName := strings.TrimPrefix(pi.ImportPath, LOCAL_PACKAGE_PATH+SPILT)
		rs.Nodes = append(rs.Nodes, data{cNode{
			idx, nodeName,
		}})
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
				rs.Links = append(rs.Links, data{cLink{nodeID, depID}})
			} else {
				rs.Nodes = append(rs.Nodes, data{cNode{
					idx, depName,
				}})
				rs.Links = append(rs.Links, data{cLink{nodeID, idx}})
				nodeMap[depName] = idx
				idx++
			}
		}
	}

	return rs
}

// loadETemplate 加载页面模板
func loadCTemplate() ([2][]byte, error) {
	f, err := os.Open(TEMPLATE_PATH)
	if err != nil {
		return [2][]byte{}, err
	}
	tByte, err := ioutil.ReadAll(f)
	if err != nil {
		return [2][]byte{}, err
	}
	return [2][]byte{tByte[:666], tByte[874:]}, nil
}

// TODO 展示页面未来写死在代码里，减少加载文件可能出现的异常与耗时
// generateEView 生成展示页面
func generateCView(rs cResultSet) string {
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
