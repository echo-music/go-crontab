package api

import (
	"encoding/json"
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/common"
	"github.com/echo-music/go-crontab/master/logic"
	"github.com/echo-music/go-crontab/master/model/etcd/keys"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net/http"
)

//添加任务
//ETCDCTL_API=3 etcdctl get  /cron/tasks/job1
func checkTaskSaveRequest(req *http.Request) (task *common.Task, err error) {
	task = &common.Task{}
	if req.Method != "POST" {
		err = status.Error(403, "不允许"+req.Method+"请求")
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	//参数绑定
	if err = common.BindBody(body, task); err != nil {
		return
	}
	if task.Name == "" {
		err = status.Error(404, "任务名称为必填项")
		return
	}
	if task.Command == "" {
		err = status.Error(404, "命令为必填项")
		return
	}
	if task.CronExpr == "" {
		err = status.Error(404, "定时规则为必填项")
		return
	}
	return
}

//添加定时任务
func TaskSave(resp http.ResponseWriter, req *http.Request) {

	var (
		putRes *etcdv3.PutResponse
		task   *common.Task
		err    error
	)

	//参数绑定与校验
	if task, err = checkTaskSaveRequest(req); err != nil {
		reply, _ := common.BuildResponse(401, err.Error(), "")
		_, err = resp.Write(reply)
		return
	}

	//添加任务
	key := keys.TASK_SAVE_DIR + task.Name
	taskByte, _ := json.Marshal(task)
	if putRes, err = logic.TaskSave(key, string(taskByte)); err != nil {
		reply, _ := common.BuildResponse(401, err.Error(), "")
		_, err = resp.Write(reply)
		return

	}
	reply, _ := common.BuildResponse(0, "success", putRes)
	_, err = resp.Write(reply)
	return

}
