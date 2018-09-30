package main

import (
	"testing"
)

func TestGenResultSet(t *testing.T) {
	type args struct {
		packageInfo map[string]packageInfo
	}
	tests := []struct {
		name string
		args args
		want resultSet
	}{
		{
			"test01",
			args{},
			resultSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := genResultSet(tt.args.packageInfo); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("genResultSet() = %v, want %v", got, tt.want)
			//}
			genResultSet(tt.args.packageInfo)
		})
	}
}
