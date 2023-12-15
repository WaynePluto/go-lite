package utils

import (
	"reflect"
	"testing"
)

func Test_struct2map(t *testing.T) {
	type Param0 struct {
		Phone     string `json:"phone" type:"string"`
		Count     int    `json:"count" type:"number"`
		IsDeleted bool   `json:"isDeleted" type:"boolean"`
	}
	type Param1 struct {
		Name string `json:"name" type:"string"`
		Age  int    `json:"age" type:"number"`
	}
	type args struct {
		param any
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
		{
			name: "test Param0",
			args: args{
				param: Param0{},
			},
			want: map[string]string{"phone": "string", "count": "number", "isDeleted": "boolean"},
		},
		{
			name: "test Param1",
			args: args{
				param: Param1{},
			},
			want: map[string]string{"name": "string", "age": "number"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("struct2map() = %v, want %v", got, tt.want)
			}
		})
	}
}
