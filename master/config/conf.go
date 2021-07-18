package config

import "time"

//定义etcd
type EtcdConfig struct {
	// Endpoints is a list of URLs.
	Endpoints   []string      `json:"endpoints"`
	DialTimeout time.Duration `json:"dial-timeout"`
}

//etcd配置初始化
var EtcdConf = &EtcdConfig{
	Endpoints:   []string{"127.0.0.1:2379"},
	DialTimeout: 5 * time.Second,
}

//mongo配置初始化
type MongoConfig struct {
	Url string `json:"url"`
}

//mongo 配置初始化
var MongoConf = &MongoConfig{
	Url: "mongodb://localhost:27017",
}
