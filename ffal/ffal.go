package ffal

import (
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(c *Context)

// 实现handler接口
type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

// 新增路由
func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method,pattern,handler)
}

// GET请求
func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.router.addRoute(GET, pattern, handlerFunc)
}

// POST请求
func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	engine.router.addRoute(POST, pattern, handlerFunc)
}

// run方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}
