package ffal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// context 对上下文的封装
type Context struct {
	// request和response-writer
	Writer http.ResponseWriter
	Req *http.Request
	// 请求路径和请求方式
	Path string
	Method string
	// 返回状态码
	StatusCode int
}

func NewContext(writer http.ResponseWriter, r *http.Request)*Context{
	return &Context{
		Writer:     writer,
		Req:        r,
		Path:       r.URL.Path,
		Method:     r.Method,
	}
}

// 返回 post form的value
func (c *Context)PostForm(key string)string{
	return c.Req.FormValue(key)
}

func (c *Context)Query(key string)string{
	return c.Req.URL.Query().Get(key)
}

func (c *Context)Status(status int){
	c.StatusCode = status
	c.Writer.WriteHeader(status)
}

func (c *Context)SetHeader(key,value string){
	c.Req.Header.Set(key,value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}