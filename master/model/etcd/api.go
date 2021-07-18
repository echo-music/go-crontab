package etcd

import (
	"context"
	etcdv3 "github.com/coreos/etcd/clientv3"
)

var (
	client *etcdv3.Client
)

//保存
func Save(key string, val string) (putRes *etcdv3.PutResponse, err error) {
	client = EtcdClient()
	kv := etcdv3.NewKV(client)
	if putRes, err = kv.Put(context.TODO(), key, val); err != nil {
		return
	}
	return
}

func Delete(key string) (delRes *etcdv3.DeleteResponse, err error) {
	client = EtcdClient()
	kv := etcdv3.NewKV(client)
	if delRes, err = kv.Delete(context.TODO(), key); err != nil {
		return
	}
	return
}
