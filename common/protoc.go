package common

import (
	"encoding/json"
	"github.com/echo-music/go-crontab/work/model/etcd"
	"github.com/gorhill/cronexpr"
	"strings"
	"time"
)

// 定时任务
type Task struct {
	Name     string `json:"name"`     //  任务名
	Command  string `json:"command"`  // shell命令
	CronExpr string `json:"cronExpr"` // cron表达式
}

// 任务调度计划
type TaskSchedulePlan struct {
	Job      *Task                // 要调度的任务信息
	Expr     *cronexpr.Expression // 解析好的cronexpr表达式
	NextTime time.Time            // 下次调度时间
}

// HTTP接口应答
type Response struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

// 应答方法
func BuildResponse(errno int, msg string, data interface{}) (resp []byte, err error) {
	// 1, 定义一个response
	var (
		response Response
	)

	response.Errno = errno
	response.Msg = msg
	response.Data = data

	// 2, 序列化json
	resp, err = json.Marshal(response)
	return
}

func BindBody(body []byte, target interface{}) (err error) {
	err = json.Unmarshal(body, target)
	return
}

// 从etcd的key中提取任务名
// /cron/jobs/job10抹掉/cron/jobs/
func ExtractJobName(jobKey string) string {
	return strings.TrimPrefix(jobKey, etcd.TASK_SAVE_DIR)
}

// 从 /cron/killer/job10提取job10
func ExtractKillerName(killerKey string) string {
	return strings.TrimPrefix(killerKey, etcd.TASK_KILLER_DIR)
}

// 任务变化事件有2种：1）更新任务 2）删除任务
func BuildJobEvent(eventType int, task *Task) (jobEvent *TaskEvent) {
	return &TaskEvent{
		EventType: eventType,
		Task:      task,
	}
}

// 变化事件
type TaskEvent struct {
	EventType int //  SAVE, DELETE
	Task      *Task
}
