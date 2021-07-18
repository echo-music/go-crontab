package logic

import (
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/master/model/etcd"
)

func TaskKill(key string, val string, ttl int64) (putReps *etcdv3.PutResponse, err error) {

	putReps, err = etcd.Kill(key, val, ttl)
	return

}
