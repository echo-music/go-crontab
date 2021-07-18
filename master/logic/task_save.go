package logic

import (
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/master/model/etcd"
)

func TaskSave(key, val string) (putRes *etcdv3.PutResponse, err error) {
	if putRes, err = etcd.Save(key, val); err != nil {
		return
	}
	return
}
