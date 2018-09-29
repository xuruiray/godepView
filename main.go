package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var GOPATH string
var WORKSPACE string
var TEMPLATE string

func main() {
	fileList := make([]string, len(os.Args))
	for idx, args := range os.Args {
		fileList[idx] = args
	}
	fileList = fileList[1:]
	fileList = paramsFilter(fileList)

	resultMap := loadData(fileList)
	for _, v := range resultMap {
		fmt.Println(v)
	}
}

type packageInfo struct {
	// 当前包自身路径
	ImportPath string
	// 当前包包含的文件
	GoFiles []string
	// 当前包的引用
	Imports []string
}

// 初始化 GOPATH 与当前路径
func init() {

	GOPATH = os.Getenv("GOPATH")

	cmd := exec.Command("pwd")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("command 'pwd' execution failed error: %v", err)
		os.Exit(0)
	}
	WORKSPACE = strings.Trim(out.String(), "\n")

	TEMPLATE, err = loadTemplate()
	if err != nil {
		log.Fatalf("load view template error: %v", err)
		os.Exit(0)
	}
}

// loadData 加载所有 package 信息
func loadData(paths []string) map[string]packageInfo {

	resultMap := make(map[string]packageInfo)

	for _, path := range paths {
		realPath := getPackagePath(path)
		p, err := getPackageInfo(realPath)
		if err == nil {
			resultMap[p.ImportPath] = *p
		}
	}

	return resultMap
}

// getPackageInfo 调用 go list 获取包结构
func getPackageInfo(packagePath string) (*packageInfo, error) {
	// packagePath 只能为包名
	out, err := exec.Command("go", "list", "-json", packagePath).Output()
	if err != nil {
		log.Fatalf("go command list command: %v error: %v",
			"go list -json "+packagePath, err)
		return nil, err
	}

	var m packageInfo
	dec := json.NewDecoder(bytes.NewReader(out))
	for {
		if err := dec.Decode(&m); err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("reading go list output: %v", err)
			return nil, err
		}
	}
	return &m, nil
}

// loadTemplate 加载页面模板
func loadTemplate() (string, error) {
	return "", nil
}

// generateView 生成展示页面
func generateView() {

}
