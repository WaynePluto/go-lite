package lite

import (
	"fmt"
	"net/http"
)

type Engine struct {
	router *Router
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) Use(pattern string, handler HandlerFunc) {
	matchData, ok := engine.router.root.match(pattern)

	if ok {
		matchData.node.addMiddleware(handler)
	} else {
		node, err := engine.router.root.insert(pattern)
		if err != nil {
			return
		}
		node.addMiddleware(handler)
	}
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	fmt.Printf("server start at: %v\n", addr)
	return http.ListenAndServe(addr, engine)
}
