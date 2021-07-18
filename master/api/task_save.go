package api

import (
	"github.com/echo-music/go-crontab/master/common"
	"net/http"
)

func TaskSave(resp http.ResponseWriter, req *http.Request) {

	reply, _ := common.BuildResponse(0, "success", "")

	resp.Write(reply)
}
