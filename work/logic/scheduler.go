package logic

import (
	"github.com/echo-music/go-crontab/common"
	"github.com/echo-music/go-crontab/work/model/etcd"
	"sync"
)

// 任务调度
type Scheduler struct {
	taskEventChan      chan *common.TaskEvent              // etcd任务事件队列
	taskPlanTable      map[string]*common.TaskSchedulePlan // 任务调度计划表
	taskExecutingTable map[string]*common.TaskExecuteInfo  // 任务执行表
	taskResultChan     chan *common.JobExecuteResult       // 任务结果队列
}

var once sync.Once
var GlobalScheduler *Scheduler

//初始化调度器
func init() {
	once.Do(func() {
		GlobalScheduler = &Scheduler{
			taskEventChan:      make(chan *common.TaskEvent, 1000),
			taskPlanTable:      make(map[string]*common.TaskSchedulePlan),
			taskExecutingTable: nil,
			taskResultChan:     nil,
		}
	})

	go GlobalScheduler.SchedulerLoop()
}

//推送事件给调度器的事件队列
func (scheduler *Scheduler) PushTaskEvent(taskEevent *common.TaskEvent) {

	GlobalScheduler.taskEventChan <- taskEevent
}

//调度器监听事件的到达
func (scheduler *Scheduler) SchedulerLoop() {
	for {
		select {
		case taskEvent := <-scheduler.taskEventChan:
			GlobalScheduler.HandlerTaskEvent(taskEvent)
		}

	}
}

//处理任务事件
func (scheduler *Scheduler) HandlerTaskEvent(event *common.TaskEvent) (err error) {

	var (
		taskSchedulePlan *common.TaskSchedulePlan
	)
	switch event.EventType {
	case etcd.TASK_EVENT_SAVE:
		//加入任务调度计划表
		if taskSchedulePlan, err = common.BuildTaskSchedulePlan(event.Task); err != nil {
			return
		}
		GlobalScheduler.taskPlanTable[event.Task.Name] = taskSchedulePlan

	case etcd.TASK_EVENT_DELETE:
		//将任务从调度计划表删除
		if _, ok := GlobalScheduler.taskPlanTable[event.Task.Name]; ok {
			delete(GlobalScheduler.taskPlanTable, event.Task.Name)
		}

	}
	return
}
