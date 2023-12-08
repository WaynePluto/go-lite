package core

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	Path   string
	Method string
	Params map[string]string
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		Params: make(map[string]string, 1),
	}
}

func (c *Context) json(code int, data any, msg string) {
	c.Writer.Header().Set("Content-Type", "application/json")
	resq := Response{code, data, msg}
	res, err := json.Marshal(resq)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
		return
	}
	c.Writer.WriteHeader(200)
	c.Writer.Write(res)
}

func (c *Context) JSON(data any) {
	c.json(200, data, "success")
}

func (c *Context) Error(code int, data string) {
	c.json(code, data, "error")
}

func (c *Context) GetReqBody() (any, error) {
	var body any
	err := json.NewDecoder(c.Req.Body).Decode(&body)
	if err != nil {
		log.Printf("Route %4s - %s", "get body error: ", err.Error())
		c.json(400, "The body json format is incorrect", "error")
		return nil, err
	}
	return body, nil
}
