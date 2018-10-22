package main

import (
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

// genCResultSet 生成 echarts 需要的结构体格式
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
			if strings.HasPrefix(dep, LOCAL_PACKAGE_PATH+SPILT+"vendor") {
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
