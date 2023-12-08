package lite

import (
	"log"
	"testing"
)

func dfs(n *Node, fn func(n *Node)) {
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
	dfs(root, func(n *Node) {
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

	_, ok := root.match("/")
	if !ok {
		t.Error("match root error")
		return
	}

	data, ok := root.match("/user/1")
	if !ok {
		t.Error("match root error")
		return
	}
	log.Printf("match /user/1 pattern: %v, params: %v", data.node.pattern, data.params)

	data, ok = root.match("/company/123/user/456")
	if !ok {
		t.Error("match root error")
		return
	}
	log.Printf("match /company/123/user/456 pattern: %v, params: %v", data.node.pattern, data.params)
}
