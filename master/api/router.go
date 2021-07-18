package api

import "net/http"

//注册路由
func RegisterRouter() (mux *http.ServeMux) {
	// 配置路由
	mux = http.NewServeMux()

	mux.HandleFunc("/task/save", TaskSave)
	mux.HandleFunc("/task/del", TaskDel)
	mux.HandleFunc("/task/list", TaskList)
	mux.HandleFunc("/task/kill", TaskKill)


	return mux

}
