package api

import "net/http"

//注册路由
func RegisterRouter() (mux *http.ServeMux) {
	// 配置路由
	mux = http.NewServeMux()

	mux.HandleFunc("/task/create", TaskCreate)
	mux.HandleFunc("/task/kill", TaskKill)
	mux.HandleFunc("/task/list", TaskList)
	mux.HandleFunc("/task/update", TaskList)

	return mux

}
