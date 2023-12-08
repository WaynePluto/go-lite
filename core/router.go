package core

import (
	"log"
	"strings"
)

type HandlerFunc func(*Context)

type node struct {
	pattern  string
	part     string
	children map[string]*node
	param    string
}

func creatRoot() *node {
	return &node{
		pattern:  "/",
		part:     "",
		children: make(map[string]*node, 1),
		param:    "",
	}
}

func (root *node) insert(path string) {
	if path == "/" {
		creatRoot()
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

		key, param := "", ""

		if part[0] == ':' {
			param = part[1:]
		} else {
			key = part
		}
		child, ok := curNode.children[key]

		if !ok {
			child = &node{
				pattern:  "",
				part:     key,
				children: make(map[string]*node, 1),
				param:    param,
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
			} else {
				log.Printf("路由：%s已注册", path)
			}
		}
		curNode = child
		i++
	}
}

func (root *node) match(path string, params map[string]string) (string, bool) {
	if path == "/" {
		return "/", true
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

		child, ok := curNode.children[part]

		if !ok {
			// 未匹配到静态节点
			child, ok = curNode.children[""]
			if ok {
				// 匹配到路由参数
				if params != nil {
					params[child.param] = part
				}
			} else {
				return "", false
			}
		}

		if ok && i == len(parts)-1 {
			return child.pattern, true
		}

		curNode = child
		i++
	}
	return "", false
}

type router struct {
	root     *node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		root:     creatRoot(),
		handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	r.root.insert(pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	pattern, ok := r.root.match(c.Path, c.Params)
	if ok {
		key := c.Method + "-" + pattern
		handler, ok := r.handlers[key]
		if ok {
			handler(c)
		} else {
			c.json(404, "404 NOT FOUND METHOD", "error")
		}
	} else {
		c.json(404, "404 NOT FOUND PATH", "error")
	}
}
