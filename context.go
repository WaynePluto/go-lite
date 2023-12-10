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
	Writer http.ResponseWriter
	Req    *http.Request
	Path   string
	Method string
	Params map[string]string
	// middleware
	handlers []HandlerFunc
	// middleware index
	index int
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

func (c *Context) json(code int, data any) {
	resq := Response{code, data}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	json.NewEncoder(c.Writer).Encode(resq)
}

func (c *Context) JSON(data any) {
	c.json(http.StatusOK, data)
}

func (c *Context) Error(code int, data string) {
	c.json(code, data)
}

func (c *Context) GetReqBody() (any, error) {
	var body any
	err := json.NewDecoder(c.Req.Body).Decode(&body)
	if err != nil {
		log.Printf("Route %4s - %s", "get body error: ", err.Error())
		c.json(http.StatusBadRequest, "The body json format is incorrect")
		return nil, err
	}
	return body, nil
}

func (c *Context) Next() {
	c.index++
	if c.handlers[c.index] != nil {
		c.handlers[c.index](c)
	}
}
