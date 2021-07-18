package etcd

import (
	etcdv3 "github.com/coreos/etcd/clientv3"
	"github.com/echo-music/go-crontab/master/config"
	"sync"
)

var (
	etcdClient *etcdv3.Client
	once       sync.Once
	err        error
)

func init() {

	once.Do(func() {
		conf := etcdv3.Config{
			Endpoints:   config.EtcdConf.Endpoints,
			DialTimeout: config.EtcdConf.DialTimeout,
		}
		if etcdClient, err = etcdv3.New(conf); err != nil {
			panic(err)
		}
	})

}

func EtcdClient() *etcdv3.Client{
	return etcdClient
}
