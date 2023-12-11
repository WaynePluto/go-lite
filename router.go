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
	querys      map[string][]string
}

func newRouter() *Router {
	return &Router{
		root:        &Node{name: "/", path: "/"},
		handlers:    map[string]HandlerFunc{},
		middlewares: map[string][]HandlerFunc{},
		querys:      map[string][]string{},
	}
}

func (r *Router) Use(pattern string, handlers ...HandlerFunc) {
	r.root.insert(pattern)
	r.middlewares[pattern] = append(r.middlewares[pattern], handlers...)
}

func (r *Router) addQuery(method string, pattern string, querys []string) {
	key := method + "-" + pattern
	r.querys[key] = querys
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
			ctx.index = -1
			ctx.Next()
		} else {
			ctx.Json(404, "404 NOT FOUND METHOD")
		}
	} else {
		ctx.Json(404, "404 NOT FOUND PATH")
	}
}
