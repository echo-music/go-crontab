package logic

import (
	"context"
	"encoding/json"
	"fmt"
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/echo-music/go-crontab/common"
	"github.com/echo-music/go-crontab/work/model/etcd"
)

func Watcher() {
	var (
		client *etcdv3.Client
		//lease         etcdv3.Lease
		//leaseReply    *etcdv3.LeaseGrantResponse
		//leaseId       etcdv3.LeaseID
		//putReps       *etcdv3.PutResponse
		getReps *etcdv3.GetResponse
		//keepAliveChan <-chan *etcdv3.LeaseKeepAliveResponse
		err error
	)

	client = etcd.EtcdClient()
	kv := etcdv3.NewKV(client)

	if getReps, err = kv.Get(context.TODO(), etcd.TASK_SAVE_DIR, etcdv3.WithPrefix()); err != nil {

		return
	}

	if len(getReps.Kvs) == 0 {
		return
	}

	//启动的时候初始化所有的任务
	for _, kvTask := range getReps.Kvs {
		task := &common.Task{}
		if err = json.Unmarshal(kvTask.Value, task); err != nil {
			continue
		}
		//调度任务
	    taskEvent :=common.BuildJobEvent(etcd.TASK_EVENT_SAVE,task)
	    fmt.Println(taskEvent)

	}

	//监听新的任务变化
	go func() {

		//从这个版本监听
		watchCurVersion := getReps.Header.Revision + 1

		//创建监听器
		watcher := etcdv3.NewWatcher(client)

		watchChan := watcher.Watch(context.TODO(), etcd.TASK_SAVE_DIR, etcdv3.WithRev(watchCurVersion),etcdv3.WithPrefix())

		for respChan := range watchChan {
			for _, event := range respChan.Events {
				switch event.Type {
				case mvccpb.PUT:
					task := &common.Task{}
					if err := json.Unmarshal(event.Kv.Value, task); err != nil {
						continue
					}
					//构建一个event事件推送给调度器
					taskEvent := common.BuildJobEvent(etcd.TASK_EVENT_SAVE, task)
					fmt.Println("修改为:", string(event.Kv.Value), "version", event.Kv.CreateRevision, event.Kv.ModRevision)
					fmt.Println(taskEvent)

				case mvccpb.DELETE:
					//删除任务
					taskName := common.ExtractJobName(string(event.Kv.Key))
					task := &common.Task{
						Name: taskName,
					}
					taskEvent := common.BuildJobEvent(etcd.TASK_EVENT_DELETE, task)
					fmt.Println("删除:", string(event.Kv.Key), "version", event.Kv.CreateRevision, event.Kv.ModRevision)
					fmt.Println(taskEvent)

				}
			}
		}

	}()


}
