package utils

import (
	"reflect"
	"testing"
)

func Test_SplitPathToParts(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"splite /user/:id",
			args{path: "/user/:id"},
			[]string{"/", "user", ":id"},
		},
		{
			"splite user/id",
			args{path: "/user/id"},
			[]string{"/", "user", "id"},
		},
		{
			"splite root path",
			args{path: "/"},
			[]string{"/"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPathToParts(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitPathToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPathToPaths(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			"splite /user/:id",
			args{path: "/user/:id"},
			[]string{"/", "/user", "/user/:id"},
		},
		{
			"splite user/id",
			args{path: "/user/id"},
			[]string{"/", "/user", "/user/id"},
		},
		{
			"splite root path",
			args{path: "/"},
			[]string{"/"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitPathToPaths(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitPathToPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
