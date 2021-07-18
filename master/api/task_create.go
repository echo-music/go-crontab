package api

import "net/http"

func TaskCreate(resp http.ResponseWriter, req *http.Request) {

	resp.Write([]byte("hello"))
}

