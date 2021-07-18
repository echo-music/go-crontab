package api

import (
	"github.com/echo-music/go-crontab/common"
	"github.com/echo-music/go-crontab/master/logic"
	"github.com/echo-music/go-crontab/master/model/etcd/keys"
	"net/http"
)

func TaskList(resp http.ResponseWriter, req *http.Request) {

	res, err := logic.TaskList(keys.TASK_SAVE_DIR)
	if err != nil {
		reply, _ := common.BuildResponse(401, err.Error(), res)
		_, err = resp.Write(reply)
		return
	}

	reply, _ := common.BuildResponse(0,  "success",res)
	_, err = resp.Write(reply)
	return
}
