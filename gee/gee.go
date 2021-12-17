package gee

import "net/http"

// HandlerFunc 定义请求处理程序
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Engine Engine需要实现ServeHTTP接口
type Engine struct {
	router map[string]HandlerFunc
}

// New 构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func ()  {
	
}