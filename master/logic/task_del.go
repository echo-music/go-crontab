package logic

import (
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/master/model/etcd"
)

func TaskDel(key string) (delRes *etcdv3.DeleteResponse, err error) {

	delRes, err = etcd.Delete(key)
	return

}
