/*
 * @Author: wtf
 * @Date: 2020-08-19 13:36:50
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-19 20:27:26
 * @Description: plase write Description
 */
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
	//base
	RunMode string
	//app
 	PageSize int
 	JwtSecret string
	//server
 	HTTPPort int
 	HTTPHost string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
)


//初始化配置
func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	Cfg = Cfg

	LoadBase()
	LoadServer()
	LoadApp()
}

//加载base配置
func LoadBase() {
	// RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	sec, err := Cfg.GetSection("base")
	if err != nil {
		log.Fatalf("加载base配置失败:%v", err)
	}
	RunMode = sec.Key("RUN_MODE").MustString("debug")
}

//加载app配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("加载app配置失败:%v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("xxxxxx@xxxaaqwe123")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("加载server配置失败:%v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8080)
	HTTPHost = sec.Key("HTTP_HOST").MustString("0.0.0.0")
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
	
}



