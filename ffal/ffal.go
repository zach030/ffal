package ffal

import (
	"net/http"
	"strings"
)

const (
	GET  = "GET"
	POST = "POST"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(c *Context)

// 实现handler接口
type Engine struct {
	*RouterGroup          // engine itself : main group
	router *router
	groups []*RouterGroup // store all groups
}

type RouterGroup struct {
	prefix      string        // prefix for a router-group
	middlewares []HandlerFunc // middlewares for current router group
	parent      *RouterGroup
	engine      *Engine
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	mainGroup := &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{mainGroup}
	return engine
}

// 新增路由
func (group *RouterGroup) addRoute(method, pattern string, handler HandlerFunc) {
	group.engine.router.addRoute(method, pattern, handler)
}

// GET请求
func (group *RouterGroup) GET(pattern string, handlerFunc HandlerFunc) {
	group.addRoute(GET, pattern, handlerFunc)
}

// POST请求
func (group *RouterGroup) POST(pattern string, handlerFunc HandlerFunc) {
	group.addRoute(POST, pattern, handlerFunc)
}

// run方法
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		// 遍历全部的group，找到当前url对应的组内的全部中间件，放入context中线性组织起来
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, req)
	// 将全部的middleware 放到context中
	c.handlers = middlewares
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

// 使用中间件
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
