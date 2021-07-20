package logic

import (
	"github.com/echo-music/go-crontab/common"
	"sync"
)

// 任务调度
type Scheduler struct {
	taskEventChan chan *common.TaskEvent	//  etcd任务事件队列
	taskPlanTable map[string]*common.TaskSchedulePlan // 任务调度计划表
	taskExecutingTable map[string]*common.TaskExecuteInfo // 任务执行表
	taskResultChan chan *common.JobExecuteResult	// 任务结果队列
}

var once  sync.Once
var _scheduler *Scheduler
//初始化调度器

func init()  {
	once.Do(func() {
		_scheduler = &Scheduler{
			taskEventChan:      make(chan *common.TaskEvent,1000),
			taskPlanTable:      nil,
			taskExecutingTable: nil,
			taskResultChan:     nil,
		}
	})
}


//推送事件给调度器