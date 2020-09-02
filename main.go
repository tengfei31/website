/*
 * @Author: wtf
 * @Date: 2020-08-18 19:10:20
 * @LastEditors: wtf
 * @LastEditTime: 2020-09-02 11:57:59
 * @Description: plase write Description
 */
package main

import (
	"flag"
	"fmt"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/robfig/cron"
	"github.com/tengfei31/website/models"
	"github.com/tengfei31/website/pkg/gredis"
	"github.com/tengfei31/website/pkg/logging"
	"github.com/tengfei31/website/pkg/setting"
	"github.com/tengfei31/website/routers"
)

var cronFlag int
var httpPort int
var httpHost string

func init() {
	flag.IntVar(&cronFlag, "cron", 0, "是否要启用cron定时任务")
	flag.IntVar(&httpPort, "port", 8088, "http server port. default -port=8088")
	flag.StringVar(&httpHost, "host", "127.0.0.1", "http server host. default -host=127.0.0.1")
}
//定时任务
func cronTask() {
	logging.Info("cron starting...")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		logging.Info("run models.CleanAllArticle...")
		models.CleanAllArticle()
	})
	c.AddFunc("* * * * * *", func() {
		logging.Info("run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}

//http server
func httpServer() {
	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf("%s:%d", setting.ServerSetting.HttpHost, setting.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}
	err := server.ListenAndServe()
	if err != nil {
		logging.Error(fmt.Sprintf("Server err: %v", err))
	}
}

func main() {
	//解析命令行参数
	flag.Parse()
	if cronFlag > 1 {
		fmt.Println("请输入-cron=1或-cron=0")
		return 
	}
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
	if cronFlag == 0 {
		fmt.Println("启用http server")
		if httpHost != "" {
			setting.ServerSetting.HttpHost = httpHost
		}
		if httpPort > 0 {
			setting.ServerSetting.HttpPort = httpPort
		}
		fmt.Printf("开启服务:%s:%d\n", setting.ServerSetting.HttpHost, setting.ServerSetting.HttpPort)
		httpServer()
	} else {
		fmt.Println("启用cron定时任务")
		cronTask()
	}
}
