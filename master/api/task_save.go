package api

import (
	"encoding/json"
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/common"
	"github.com/echo-music/go-crontab/master/logic"
	"github.com/echo-music/go-crontab/master/model/etcd/keys"
	"google.golang.org/grpc/status"
	"net/http"
)

//添加任务

func checkTaskSaveRequest(task common.Task) (err error) {
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
func TaskSave(resp http.ResponseWriter, req *http.Request) {

	var (
		putRes *etcdv3.PutResponse
		err    error
	)

	if req.Method != "get" {
		reply, _ := common.BuildResponse(401, "拒绝请求", "")
		_, err = resp.Write(reply)
		return
	}
	req.GetBody
	// 1, 解析POST表单
	if err = req.ParseForm(); err != nil {
		reply, _ := common.BuildResponse(401, "表单有问题", "")
		_, err = resp.Write(reply)
		return
	}
	// 2, 取表单中的job字段
	postTask := req.PostForm.Get("job")
	task := &common.Task{}
	// 3, 反序列化job
	if err = json.Unmarshal([]byte(postTask), &task); err != nil {
		reply, _ := common.BuildResponse(401, err.Error(), "")
		_, err = resp.Write(reply)
		return
	}

	//添加任务
	key := keys.TASK_SAVE_DIR + task.Name
	taskByte, err := json.Marshal(task)
	if putRes, err = logic.TaskSave(key, string(taskByte)); err != nil {
		reply, _ := common.BuildResponse(401, err.Error(), "")
		_, err = resp.Write(reply)
		return

	} else {
		reply, _ := common.BuildResponse(0, "success", putRes)
		_, err = resp.Write(reply)
		return
	}

}
