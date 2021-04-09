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
	*RouterGroup
	router *router
	groups []*RouterGroup // store all groups
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 新增路由
func (group *RouterGroup) addRoute(method, pattern string, handler HandlerFunc) {
	group.engine.router.addRoute(method, pattern, handler)
}

// GET请求
func (group *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	group.engine.router.addRoute(GET, pattern, handlerFunc)
}

// POST请求
func (group *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	group.engine.router.addRoute(POST, pattern, handlerFunc)
}

// run方法
func (group *RouterGroup) Run(addr string) (err error) {
	return http.ListenAndServe(addr, group.engine)
}

// 实现接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}

// 新增路由组，在当前路由组下新建一个，注意需要设置parent，添加到engine里的list中
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix:      prefix,
		middlewares: nil,
		parent:      group,
		engine:      group.engine,
	}
	group.engine.groups = append(group.engine.groups, newGroup)
	return newGroup
}

type RouterGroup struct {
	prefix      string        // prefix for a router-group
	middlewares []HandlerFunc // middlewares for current router group
	parent      *RouterGroup
	engine      *Engine
}
