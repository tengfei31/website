/*
 * @Author: wtf
 * @Date: 2020-08-19 13:36:50
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 19:33:40
 * @Description: plase write Description
 */
package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string 
	PageSize int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath string
	ImageMaxSize int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}
var AppSetting = &App{}

type Server struct {
	RunMode string
	HttpHost string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &Server{}

type DataBase struct {
	Type string
	Host string
	User string
	Password string
	DbName string
	TablePrefix string
}
var DataBaseSetting  = &DataBase{}

type Redis struct {
	Type string
	Host string
	Password string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}
var RedisSetting = &Redis{}

var DefaultMb int = 1024 * 1024

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
	AppSetting.ImageMaxSize *= DefaultMb

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	err = Cfg.Section("database").MapTo(DataBaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
	RedisSetting.IdleTimeout *= time.Second
}



