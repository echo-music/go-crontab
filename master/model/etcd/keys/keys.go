package keys

const (
	// 任务保存目录
	TASK_SAVE_DIR = "/cron/tasks/"

	// 任务强杀目录
	TASK_KILLER_DIR = "/cron/killer/"

	// 任务锁目录
	TASK_LOCK_DIR = "/cron/lock/"

	// 服务注册目录
	TASK_WORKER_DIR = "/cron/workers/"

	// 保存任务事件
	TASK_EVENT_SAVE = 1

	// 删除任务事件
	TASK_EVENT_DELETE = 2

	// 强杀任务事件
	TASK_EVENT_KILL = 3
)