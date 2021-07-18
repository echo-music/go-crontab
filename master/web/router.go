package web

import (
	"github.com/echo-music/go-crontab/master/api"
	"net/http"
)

//注册路由
func RegisterRouter() (mux *http.ServeMux) {
	// 配置路由
	mux = api.RegisterRouter()


	mux.HandleFunc("/home", Index)

	staticDir :=http.Dir("./master/views/home")


	staticHandler :=http.FileServer(staticDir)
	mux.Handle("/",http.StripPrefix("/",staticHandler))
	return mux

}