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

//删除
func Delete(key string) (delRes *etcdv3.DeleteResponse, err error) {
	client = EtcdClient()
	kv := etcdv3.NewKV(client)
	if delRes, err = kv.Delete(context.TODO(), key); err != nil {
		return
	}
	return
}

//详情
func FindOne(key string) (getRes *etcdv3.GetResponse, err error) {
	client = EtcdClient()
	kv := etcdv3.NewKV(client)
	if getRes, err = kv.Get(context.TODO(), key); err != nil {
		return
	}
	return
}

//列表
func FindAll(key string) (getRes *etcdv3.GetResponse, err error) {
	client = EtcdClient()
	kv := etcdv3.NewKV(client)
	if getRes, err = kv.Get(context.TODO(), key, etcdv3.WithPrefix()); err != nil {
		return
	}
	return
}
