package main

import (
	"fmt"
	"github.com/echo-music/go-crontab/master/model/etcd"
	"github.com/echo-music/go-crontab/master/model/mongo"
)

func main() {
	client :=etcd.EtcdClient()
	fmt.Println(client)
	client =etcd.EtcdClient()
	fmt.Println(client)
	client =etcd.EtcdClient()
	fmt.Println(client)

	mongoClient :=mongo.MongoClient()
	fmt.Println(mongoClient)
	mongoClient =mongo.MongoClient()
	fmt.Println(mongoClient)
	mongoClient =mongo.MongoClient()
	fmt.Println(mongoClient)

}
