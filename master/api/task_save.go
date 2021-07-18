package api

import (
	"context"
	"fmt"
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/master/common"
	"github.com/echo-music/go-crontab/master/model/etcd"
	"net/http"
)

func TaskSave( resp http.ResponseWriter,req *http.Request)  {

	var (
		client *etcdv3.Client
		putRes *etcdv3.PutResponse
		err error
	)

	fmt.Println(req)

	client = etcd.EtcdClient()
	kv := etcdv3.NewKV(client)

	//添加任务
	if putRes, err = kv.Put(context.TODO(), "/cron/task/task1", "hello1"); err != nil {
		reply, _ := common.BuildResponse(0, err.Error(), "")
		_, err = resp.Write(reply)


	} else {
		reply, _ := common.BuildResponse(0, "success", putRes)
		_, err = resp.Write(reply)
	}



}
