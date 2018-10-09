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
		want eResultSet
	}{
		{
			"test01",
			args{},
			eResultSet{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := genEResultSet(tt.args.packageInfo); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("genEResultSet() = %v, want %v", got, tt.want)
			//}
			genEResultSet(tt.args.packageInfo)
		})
	}
}
