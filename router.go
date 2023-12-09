package lite

import (
	"log"

	"github.com/WaynePluto/go-lite/utils"
)

type HandlerFunc func(ctx *Context)

type Router struct {
	root        *Node
	handlers    map[string]HandlerFunc
	middlewares map[string][]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		root:        &Node{name: "/", path: "/"},
		handlers:    map[string]HandlerFunc{},
		middlewares: map[string][]HandlerFunc{},
	}
}

func (r *Router) Use(pattern string, handler HandlerFunc) {
	r.root.insert(pattern)
	r.middlewares[pattern] = append(r.middlewares[pattern], handler)
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %7s - %s", method, pattern)
	r.root.insert(pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) handle(ctx *Context) {
	matchData, ok := r.root.match(ctx.Path)
	if ok {
		pattern := matchData.node.path
		key := ctx.Method + "-" + pattern
		handler, ok := r.handlers[key]
		if ok {
			middlewareKeys := utils.SplitPathToPaths(pattern)
			middlewares := []HandlerFunc{}

			for _, key := range middlewareKeys {
				if ms, ok := r.middlewares[key]; ok {
					middlewares = append(middlewares, ms...)
				}
			}

			ctx.Params = matchData.params
			ctx.handlers = append(middlewares, handler)
			ctx.index = 0
			ctx.Next()
		} else {
			ctx.json(404, "404 NOT FOUND METHOD", "error")
		}
	} else {
		ctx.json(404, "404 NOT FOUND PATH", "error")
	}
}
