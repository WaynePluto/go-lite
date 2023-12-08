package lite

import (
	"log"
)

type HandlerFunc func(ctx *Context)

type Router struct {
	root     *Node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		root:     creatRoot(),
		handlers: make(map[string]HandlerFunc)}
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	r.root.insert(pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) handle(c *Context) {
	data, ok := r.root.match(c.Path)
	if ok {
		key := c.Method + "-" + data.node.pattern
		handler, ok := r.handlers[key]
		if ok {
			c.Params = data.params
			c.handlers = data.middlewares
			c.handlers = append(c.handlers, handler)
			c.index = 0
			c.Next()
		} else {
			c.json(404, "404 NOT FOUND METHOD", "error")
		}
	} else {
		c.json(404, "404 NOT FOUND PATH", "error")
	}
}
