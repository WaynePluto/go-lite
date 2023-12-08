package lite

import (
	"errors"
	"log"
	"strings"
)

type Node struct {
	// 节点的匹配模式
	pattern string
	// 节点所在部分
	part string
	// 动态参数
	param string
	// 节点中间件
	middlewares []HandlerFunc

	children map[string]*Node
}

func creatRoot() *Node {
	return &Node{
		pattern:     "/",
		middlewares: make([]HandlerFunc, 1),
		children:    make(map[string]*Node, 1),
	}
}

func (root *Node) insert(path string) (*Node, error) {
	if path == "/" {
		return creatRoot(), nil
	}

	parts := strings.Split(path, "/")
	if path[0] == '/' {
		parts = parts[1:]
	}
	curNode := root
	i := 0
	for curNode != nil && i < len(parts) {
		part := parts[i]
		if part == "" {
			i++
			continue
		}

		key, param := "", ""

		if part[0] == ':' {
			param = part[1:]
		} else {
			key = part
		}
		child, ok := curNode.children[key]

		if !ok {
			child = &Node{
				part:        key,
				param:       param,
				middlewares: make([]HandlerFunc, 1),
				children:    make(map[string]*Node, 1),
			}
			curNode.children[key] = child
		}

		if param != child.param {
			log.Printf("动态路由参数：%s已注册为%s", param, child.param)
		}

		if i == len(parts)-1 {
			// fmt.Printf("匹配到最后一层：%s\n", child.part)
			if child.pattern == "" {
				child.pattern = path
				return child, nil
			} else {
				log.Printf("路由：%s已注册", path)
			}
		}
		curNode = child
		i++
	}
	return nil, errors.New("create node error")
}

type matchData struct {
	node        *Node
	params      map[string]string
	middlewares []HandlerFunc
}

// 路由匹配，需要匹配 出最后的节点、所有中间件、所有路径参数
func (root *Node) match(path string) (data matchData, matched bool) {
	data = matchData{params: make(map[string]string, 1)}
	matched = false
	if path == "/" {
		data.node = root
		data.middlewares = data.node.middlewares
		matched = true
		return
	}
	parts := strings.Split(path, "/")
	if path[0] == '/' {
		parts = parts[1:]
	}
	curNode := root
	i := 0
	for curNode != nil && i < len(parts) {
		part := parts[i]
		if part == "" {
			i++
			continue
		}
		// 匹配静态节点
		child, ok := curNode.children[part]
		if !ok {
			// 未匹配到静态节点，尝试匹配动态节点
			child, ok = curNode.children[""]
			if ok {
				data.params[child.param] = part
			}
		}
		if ok {
			// 匹配成功
			data.middlewares = append(data.middlewares, child.middlewares...)

			// 匹配完成
			if i == len(parts)-1 {
				matched = true
				data.node = child
				break
			}
		} else {
			// 匹配失败
			break
		}

		curNode = child
		i++
	}
	return data, matched
}

func (n *Node) addMiddleware(handler HandlerFunc) {
	if n.middlewares == nil {
		n.middlewares = make([]HandlerFunc, 1)
	}
	n.middlewares = append(n.middlewares, handler)
}
