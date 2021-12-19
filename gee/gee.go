package gee

import (
	"net/http"
)

// HandlerFunc 定义请求处理程序
type HandlerFunc func(c *Context)

// Engine Engine需要实现ServeHTTP接口
type Engine struct {
	router *router // 路由映射表
}

// 实现ServeHTTP接口
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	//key := req.Method + "-" + req.URL.Path
	//if hander, ok := engine.router[key]; ok {
	//	hander(w, req)
	//} else {
	//	w.WriteHeader(http.StatusNotFound)
	//	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	//}
	c := newContext(w, req)
	engine.router.handle(c)
}

// New 构造函数
func New() *Engine {
	return &Engine{router: newRouter()}
}

// 路由表
func (engine *Engine) addRoute(method string, pattern string, hander HandlerFunc)  {
	engine.router.addRoute(method, pattern, hander)
}

// GET 添加GET请求
func (engine *Engine) GET(pattern string, handler HandlerFunc)  {
	engine.addRoute("GET", pattern, handler)
}

// POST 添加POST请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 启动函数 ListenAndServe 的包装
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}