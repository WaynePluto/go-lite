package core

import (
	"log"
	"testing"
)

func dfs(n *node, fn func(n *node)) {
	fn(n)
	for _, v := range n.children {
		dfs(v, fn)
	}
}

func Test_insert(t *testing.T) {
	root := creatRoot()
	root.insert("/user/:id")
	root.insert("/role/user/:id")

	count := 0
	dfs(root, func(n *node) {
		log.Printf("dfs node part: %s, param:%s, pattern: %s \n", n.part, n.param, n.pattern)
		count++
	})
	if count != 6 {
		t.Error("dfs node count error", count)
	}
}

func Test_match(t *testing.T) {
	root := creatRoot()
	root.insert("/user/:id")
	root.insert("/role/user/:id")
	root.insert("/company/:cid/user/:uid")

	_, ok := root.match("/", nil)
	if !ok {
		t.Error("match root error")
		return
	}

	params := make(map[string]string, 1)
	pattern, ok := root.match("/user/1", params)
	if !ok {
		t.Error("match root error")
		return
	}
	log.Printf("match /user/1 pattern: %v, params: %v", pattern, params)

	params = make(map[string]string, 1)
	pattern, ok = root.match("/company/123/user/456", params)
	if !ok {
		t.Error("match root error")
		return
	}
	log.Printf("match /company/123/user/456 pattern: %v, params: %v", pattern, params)

}
