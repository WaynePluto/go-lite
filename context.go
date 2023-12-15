package lite

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type Context struct {
	Writer   http.ResponseWriter
	Req      *http.Request
	Path     string
	Method   string
	Params   map[string]string
	Err      error
	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		Params: make(map[string]string),
		index:  -1,
	}
}

func (c *Context) Json(code int, data any) {
	resq := Response{code, data}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(200)
	json.NewEncoder(c.Writer).Encode(resq)
}

func (c *Context) JSON(data any) {
	c.Json(http.StatusOK, data)
}

func (c *Context) Query() map[string][]string {
	return c.Req.URL.Query()
}

// 获取请求参数，不需要再次进行错误处理
func (c *Context) Body() (any, bool) {
	var body any
	err := json.NewDecoder(c.Req.Body).Decode(&body)
	if err != nil {
		log.Printf("Route %4s - %s", "get body error: ", err.Error())
		c.Json(http.StatusBadRequest, "The body json format is incorrect")
		return nil, false
	}
	return body, true
}

func (c *Context) Next() {
	c.index++
	if c.handlers[c.index] != nil {
		c.handlers[c.index](c)
	}
}
