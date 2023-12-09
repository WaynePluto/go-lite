package lite

import (
	"log"
	"net/http"
)

type Engine struct {
	router *Router
}

type ApiParam struct {
	query []string
	body  any
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc, query []string) {
	engine.router.addRoute(method, pattern, handler)
	engine.router.addQuery(method, pattern, query)
}

func (engine *Engine) Use(pattern string, handlers ...HandlerFunc) {
	engine.router.Use(pattern, handlers...)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc, query []string) {
	engine.addRoute("GET", pattern, handler, query)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler, nil)
}

func (engine *Engine) PUT(pattern string, handler HandlerFunc) {
	engine.addRoute("PUT", pattern, handler, nil)
}

func (engine *Engine) PATCH(pattern string, handler HandlerFunc) {
	engine.addRoute("PATCH", pattern, handler, nil)
}

func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute("DELETE", pattern, handler, nil)
}

func (engine *Engine) Run(addr string) error {
	log.Printf("server start at: %v\n", addr)
	return http.ListenAndServe(addr, engine)
}
