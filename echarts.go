package main

import (
	"strings"
)

type eResultSet struct {
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

// genEResultSet 生成 echarts 需要的结构体格式
func genEResultSet(packageInfoMap map[string]packageInfo) eResultSet {

	rs := eResultSet{
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
			if strings.HasPrefix(dep, LOCAL_PACKAGE_PATH+SPILT+"vendor") {
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
