// TODO 循环中删除数组元素 使用 range 的坑

package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuruiray/godepView/template"
	"os"
	"strings"
	"sync"
)

const (
	SPILT = "/"
)

// genPackageList 过滤不合法路径，补充文件夹内子路径
func genPackageList(paths []string) (resultList []string) {

	paths = addChildDir(paths)
	paths = delFilePath(paths)

	for _, path := range paths {
		if isGoPackage(path) {
			resultList = append(resultList, path)
		}
	}

	return resultList
}

// delFilePath 删除文件类型的路径 只保留文件夹类型路径
func delFilePath(paths []string) (resultList []string) {

	for _, path := range paths {

		if fi, err := os.Stat(path); err == nil && fi.IsDir() {
			resultList = append(resultList, path)
		}
	}

	return resultList
}

// addChildDir 获取文件夹内所有子文件
func addChildDir(paths []string) []string {
	var childPath []string
	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 跳过隐藏文件
			// FIXME：这段代码在并发情况下会出现误判情况
			nameArray := strings.Split(path, SPILT)
			if strings.HasPrefix(nameArray[len(nameArray)-1], ".") {
				return
			}

			// 跳过非文件夹类型文件
			if fi, err := os.Stat(path); err != nil || !fi.IsDir() {
				return
			}

			f, err := os.Open(path)
			if err != nil {
				return
			}

			names, err := f.Readdirnames(0)
			if err != nil {
				return
			}
			f.Close()

			for _, name := range names {
				childPath = append(childPath, path+SPILT+name)
			}
		}()
	}

	wg.Wait()
	if len(childPath) > 0 {
		childPath = addChildDir(childPath)
		paths = append(paths, childPath...)
	}

	return paths
}

// isGoPackage 检查文件夹是否为 go package
func isGoPackage(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}

	names, err := f.Readdirnames(0)
	if err != nil {
		return false
	}
	f.Close()

	if len(names) > 0 {
		for _, n := range names {
			if strings.HasSuffix(n, ".go") {
				return true
			}
		}
	}

	return false
}

// getPackagePath 将文件夹路径转化为 go package 路径
func getPackagePath(path string) string {

	// 相对路径转换为绝对路径
	if strings.HasPrefix(path, ".") {
		path = strings.Trim(path, ".")
		path = WORKSPACE + path
	}

	// 绝对路径转化为完整包名
	path = strings.TrimPrefix(path, GOPATH+SPILT+"src"+SPILT)
	return path
}

// generateViewFile 生成展示页面
func genFile(rs []byte, template []string) string {

	pathArray := strings.Split(LOCAL_PACKAGE_PATH, SPILT)
	fileName := pathArray[len(pathArray)-1]
	viewFilePath := WORKSPACE + SPILT + fileName + ".html"

	f, err := os.OpenFile(viewFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	defer f.Close()

	f.Write([]byte(template[0]))
	f.Write(rs)
	f.Write([]byte(template[1]))

	return viewFilePath
}

func genView(templateMode string, resultMap map[string]packageInfo) string {

	var resultByte []byte
	temp := template.GetTemplate(templateMode)

	if len(resultMap) < 8 {
		templateMode = template.CTemp
	}

	if templateMode == template.CTemp {
		resultSet := genCResultSet(resultMap)
		resultByte, _ = json.Marshal(resultSet)
	} else if templateMode == template.ETemp {
		resultSet := genCResultSet(resultMap)
		resultByte, _ = json.Marshal(resultSet)
	}

	return genFile(resultByte, temp)
}
