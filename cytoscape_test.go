package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_genResultSet(t *testing.T) {
	type args struct {
		packageInfoMap map[string]packageInfo
	}
	tests := []struct {
		name string
		args args
		want cResultSet
	}{
		{
			"test01",
			args{},
			cResultSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := genResultSet(tt.args.packageInfoMap); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("genResultSet() = %v, want %v", got, tt.want)
			//}
			got := genCResultSet(tt.args.packageInfoMap)
			result, _ := json.Marshal(got)
			fmt.Println(string(result))
		})
	}
}

func Test_loadCTemplate(t *testing.T) {
	tests := []struct {
		name    string
		want    [2][]byte
		wantErr bool
	}{
		{
			name:    "test01",
			want:    [2][]byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadCTemplate()
			if (err != nil) != tt.wantErr {
				t.Errorf("loadCTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//fmt.Println(string(got[0]))
			fmt.Println(string(got[1]))
		})
	}
}
