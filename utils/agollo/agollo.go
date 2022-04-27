package main

/*
@ function: apollo配置管理中心golang客户端演示
*/
import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/component/log"
	"github.com/apolloconfig/agollo/v4/env/config"
)

func main() {
	c := &config.AppConfig{
		AppID:             "SampleApp",
		Cluster:           "dev",
		NamespaceName:     "application",
		IP:                "http://192.168.2.102:8080",
		IsBackupConfig:    true,
		Secret:            "6ase8weg5w8s7gsg8e63g55",
		SyncServerTimeout: 0,
	}
	agollo.SetLogger(&log.DefaultLogger{})
	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	fmt.Println("初始化apollo配置成功")
	cache := client.GetConfigCache(c.NamespaceName)
	timeout, _ := cache.Get("timeout")
	honey, _ := cache.Get("honey")
	fmt.Println(timeout, honey)
}
