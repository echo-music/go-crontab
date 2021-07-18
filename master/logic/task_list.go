package logic

import (
	"encoding/json"
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/common"
	"github.com/echo-music/go-crontab/master/model/etcd"
)

//任务列表
func TaskList(key string) (taskList []*common.Task, err error) {
	var getRes *etcdv3.GetResponse
	if getRes, err = etcd.FindAll(key); err != nil {
		return
	}

	taskList = make([]*common.Task, 0)
	for _, kvPair := range getRes.Kvs {
		task := &common.Task{}
		if err = json.Unmarshal(kvPair.Value, task); err != nil {
			continue
		}
		taskList = append(taskList, task)
	}

	return
}
