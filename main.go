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
var TEMPLATE [2][]byte // 没想到还可以这样限制数组长度。。
var LOCAL_PACKAGE_PATH string
var TEMPLATE_PATH string
var C_TEMPLATE_PATH string

// 入参为包地址 ./ufs
func main() {

	LOCAL_PACKAGE_PATH = argsFilter(os.Args)[1]

	fileList := paramsFilter([]string{LOCAL_PACKAGE_PATH})
	LOCAL_PACKAGE_PATH = getPackagePath(LOCAL_PACKAGE_PATH)

	resultMap := loadData(fileList)
	rs := genCResultSet(resultMap)
	response := generateCView(rs)

	fmt.Println(response)
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

	TEMPLATE_PATH = GOPATH + "/src/github.com/xuruiray/godepView/cytoscape.html"
	TEMPLATE, err = loadCTemplate()
	if err != nil {
		log.Fatalf("load view template error: %v", err)
		os.Exit(0)
	}

}

// TODO 校验是否只有一个参数，若校验失败，程序退出，打印使用方法
// TODO 可以使用 -{arg} 的形式添加其他参数
// TODO 目前没有能力检查其他错误的原因，所以忽略其他错误，只有参数错误会退出并显示错误信息
func argsFilter(args []string) []string {
	if len(args) != 2 {
		fmt.Println("指令参数错误")
		os.Exit(0)
	}

	args[1] = strings.TrimSuffix(args[1], SPILT)
	return args
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
