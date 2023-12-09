package lite

import (
	"fmt"
	"testing"
)

func TestNode_insert(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		root *Node
		args args
		want *Node
	}{
		{
			name: "insert /user",
			root: &Node{name: "/", path: "/", children: map[string]*Node{}},
			args: args{path: "/user"},
			want: &Node{name: "user", path: "/user", paramName: ""},
		},
		{
			name: "insert /user/:id",
			root: &Node{name: "/", path: "/", children: map[string]*Node{}},
			args: args{path: "/user/:id"},
			want: &Node{name: ":", path: "/user/:id", paramName: "id"},
		},
		{
			name: "insert /user/:id/test/:cid",
			root: &Node{name: "/", path: "/", children: map[string]*Node{}},
			args: args{path: "/user/:id/test/:cid"},
			want: &Node{name: ":", path: "/user/:id/test/:cid", paramName: "cid"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.root.insert(tt.args.path)
			if err != nil {
				t.Errorf("Node.insert() error = %v \n", err)
				return
			}
			if got.path != tt.want.path || got.name != tt.want.name || got.paramName != tt.want.paramName {
				t.Errorf("Node.insert() = %v, want %v\n", got, tt.want)
				return
			}
			fmt.Printf("node value: %v\n", Node{name: got.name, path: got.path, paramName: got.paramName})
		})
	}
}

func TestNode_match(t *testing.T) {
	type args struct {
		insertPath string
		matchPath  string
	}
	tests := []struct {
		name        string
		root        *Node
		args        args
		wantData    *matchData
		wantMatched bool
	}{
		{
			name: "match /user",
			root: &Node{name: "/", path: "/", children: map[string]*Node{}},
			args: args{insertPath: "/user", matchPath: "/user"},
			wantData: &matchData{
				node:   &Node{name: "user", path: "/user", paramName: ""},
				params: map[string]string{},
			},
			wantMatched: true,
		},
		{
			name: "match /user/roles",
			root: &Node{name: "/", path: "/", children: map[string]*Node{}},
			args: args{insertPath: "/user/roles", matchPath: "/user/roles"},
			wantData: &matchData{
				node:   &Node{name: "roles", path: "/user/roles", paramName: ""},
				params: map[string]string{},
			},
			wantMatched: true,
		},
		{
			name: "match /user/:id",
			root: &Node{name: "/", path: "/", children: map[string]*Node{}},
			args: args{insertPath: "/user/:id", matchPath: "/user/1"},
			wantData: &matchData{
				node:   &Node{name: ":", path: "/user/:id", paramName: "id"},
				params: map[string]string{"id": "1"},
			},
			wantMatched: true,
		},
		{
			name: "match /user/:id/role",
			root: &Node{name: "/", path: "/", children: map[string]*Node{}},
			args: args{insertPath: "/user/:id/role", matchPath: "/user/10/role"},
			wantData: &matchData{
				node:   &Node{name: "role", path: "/user/:id/role", paramName: ""},
				params: map[string]string{"id": "10"},
			},
			wantMatched: true,
		},
		{
			name: "match /user/:id/role/:rid",
			root: &Node{name: "/", path: "/", children: map[string]*Node{}},
			args: args{insertPath: "/user/:id/role/:rid", matchPath: "/user/1/role/2"},
			wantData: &matchData{
				node:   &Node{name: ":", path: "/user/:id/role/:rid", paramName: "rid"},
				params: map[string]string{"id": "1", "rid": "2"},
			},
			wantMatched: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.root.insert(tt.args.insertPath)
			gotData, gotMatched := tt.root.match(tt.args.matchPath)
			if gotMatched != tt.wantMatched {
				t.Errorf("Node.match() matched = %v, want %v", gotMatched, tt.wantMatched)
				return
			}
			gotNode := gotData.node
			wantNode := tt.wantData.node
			if gotNode.path != wantNode.path || gotNode.name != wantNode.name || gotNode.paramName != wantNode.paramName {
				t.Errorf("Node.match() node = %v, want %v\n", gotNode, wantNode)
				return
			}

			getParams := gotData.params
			wantParams := tt.wantData.params
			for k, v := range wantParams {
				if v != getParams[k] {
					t.Errorf("Node.match() param = %v, want %v\n", getParams[k], v)
					return
				}
			}

		})
	}
}
