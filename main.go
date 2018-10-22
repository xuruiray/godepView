package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuruiray/godepView/template"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

// 全局默认环境变量
var GOPATH string
var WORKSPACE string
var LOCAL_PACKAGE_PATH string
var LOCAL_FILE_PATH string

// 用户请求
type request struct {
	GoPath           string // -g
	LocalPackagePath string // -p
	TemplateMode     string // -c
}

// 入参为包地址 ./ufs
func main() {

	// 参数解析
	req := argsFilter(os.Args)
	// 获取包内文件子包文件路径
	fileList := genPackageList([]string{LOCAL_FILE_PATH})
	// 根据文件路径获取包依赖信息
	resultMap := loadPackageInfo(fileList)
	// 生成 view 文件，返回文件路径
	response := genView(req.TemplateMode, resultMap)

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
	// 加载环境变量
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
}

// 处理请求参数
func argsFilter(args []string) request {

	// 加载 local 参数
	if len(args) != 2 {
		fmt.Println("指令参数错误")
		os.Exit(0)
	}
	args[1] = strings.TrimSuffix(args[1], SPILT)
	LOCAL_FILE_PATH = args[1]
	LOCAL_PACKAGE_PATH = getPackagePath(LOCAL_FILE_PATH)

	// 加载默认参数
	var req = request{
		GoPath:           GOPATH,
		LocalPackagePath: LOCAL_PACKAGE_PATH,
		TemplateMode:     template.CTemp,
	}

	// TODO 根据其他参数修改默认参数
	return req
}

// loadPackageInfo 加载所有 package 信息
func loadPackageInfo(paths []string) map[string]packageInfo {

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
