package lite

import (
	"log"
	"net/http"
)

type Engine struct {
	router *Router
}

// api接口参数定义
type ApiParam struct {
	Query []string
	Body  map[string]string
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc, apiParam *ApiParam) {
	engine.router.addRoute(method, pattern, handler)
	if apiParam != nil {
		if apiParam.Query != nil {
			engine.router.addQuery(method, pattern, apiParam.Query)
		}
		if apiParam.Body != nil {
			engine.router.addBody(method, pattern, apiParam.Body)
		}
	}
}

func (engine *Engine) Use(pattern string, handlers ...HandlerFunc) {
	engine.router.Use(pattern, handlers...)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc, apiParam *ApiParam) {
	engine.addRoute("GET", pattern, handler, apiParam)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc, apiParam *ApiParam) {
	engine.addRoute("POST", pattern, handler, apiParam)
}

func (engine *Engine) PUT(pattern string, handler HandlerFunc, apiParam *ApiParam) {
	engine.addRoute("PUT", pattern, handler, apiParam)
}

func (engine *Engine) PATCH(pattern string, handler HandlerFunc, apiParam *ApiParam) {
	engine.addRoute("PATCH", pattern, handler, apiParam)
}

func (engine *Engine) DELETE(pattern string, handler HandlerFunc, apiParam *ApiParam) {
	engine.addRoute("DELETE", pattern, handler, apiParam)
}

func (engine *Engine) Run(addr string) error {
	log.Printf("server start at: %v\n", addr)
	return http.ListenAndServe(addr, engine)
}
