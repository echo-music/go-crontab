package api

import (
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/common"
	"github.com/echo-music/go-crontab/master/logic"
	"github.com/echo-music/go-crontab/master/model/etcd/keys"
	"google.golang.org/grpc/status"
	"net/http"
)

//添加定时任务参数校验
//ETCDCTL_API=3 etcdctl get  /cron/tasks/job1
func checkTaskDelRequest(req *http.Request) (task *common.Task, err error) {
	task = &common.Task{}
	if req.Method != "GET" {
		err = status.Error(403, "不允许"+req.Method+"请求")
		return
	}
	query := req.URL.Query()
	task.Name = query.Get("name")

	if task.Name == "" {
		err = status.Error(404, "任务名称为必填项")
		return
	}

	return
}

//添加定时任务
func TaskDel(resp http.ResponseWriter, req *http.Request) {

	var (
		delRes *etcdv3.DeleteResponse
		task   *common.Task
		err    error
	)

	//参数绑定与校验
	if task, err = checkTaskDelRequest(req); err != nil {
		reply, _ := common.BuildResponse(401, err.Error(), "")
		_, err = resp.Write(reply)
		return
	}

	//删除任务
	key := keys.TASK_SAVE_DIR + task.Name
	if delRes, err = logic.TaskDel(key); err != nil {
		reply, _ := common.BuildResponse(401, err.Error(), "")
		_, err = resp.Write(reply)
		return
	}
	reply, _ := common.BuildResponse(0, "success", delRes)
	_, err = resp.Write(reply)
	return

}
