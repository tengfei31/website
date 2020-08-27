/*
 * @Author: wtf
 * @Date: 2020-08-18 19:10:20
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-27 17:07:23
 * @Description: plase write Description
 */
package main

import (
	"fmt"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/tengfei31/website/pkg/logging"
	"github.com/tengfei31/website/pkg/setting"
	"github.com/tengfei31/website/routers"
)

func main() {
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf("%s:%d", setting.HTTPHost, setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		logging.Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}
	err := server.ListenAndServe()
	if err != nil {
		logging.Error(fmt.Sprintf("Server err: %v", err))
	}


	// server := &http.Server{
	// 	Addr: fmt.Sprintf("%s:%d", setting.HTTPHost, setting.HTTPPort),
	// 	Handler: routers.InitRouter(),
	// 	ReadTimeout: setting.ReadTimeout,
	// 	WriteTimeout: setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// server.ListenAndServe()
}
