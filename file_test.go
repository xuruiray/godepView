// need export GOPATH=/Users/xurui/go_workspace

package main

import (
	"reflect"
	"testing"
)

func Test_getPackagePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test01",
			args{
				"./main",
			},
			"github.com/xuruiray/godepView/main",
		}, {
			"test02",
			args{
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView",
			},
			"github.com/xuruiray/godepView",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPackagePath(tt.args.path); got != tt.want {
				t.Errorf("getPackagePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_delFilePath(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test01",
			args{
				[]string{
					"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView",
					"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/main.go",
					"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/unexist",
					"./",
					"./main.go",
					"./unexist",
				},
			},
			[]string{
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView",
				"./",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delFilePath(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addChildDir(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test01",
			args{
				[]string{
					"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView",
				},
			},
			[]string{
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/LICENSE",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/template",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/file_test.go",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/README.md",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/file.go",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/.git",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/main.go",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/main_test.go",
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/.idea",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addChildDir(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addChildDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isGoPackage(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test01",
			args{
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView",
			},
			true,
		}, {
			"test02",
			args{
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/.git",
			},
			false,
		}, {
			"test03",
			args{
				"/Users/xurui/go_workspace/src/github.com/xuruiray/godepView/.idea",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGoPackage(tt.args.path); got != tt.want {
				t.Errorf("%v isGoPackage() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
