/*
 * @Author: wtf
 * @Date: 2020-08-18 19:10:20
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-20 17:26:47
 * @Description: plase write Description
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/tengfei31/website/pkg/setting"
	"github.com/tengfei31/website/routers"
)

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", setting.HTTPHost, setting.HTTPPort),
		Handler: routers.InitRouter(),
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
