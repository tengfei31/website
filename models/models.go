/*
 * @Author: wtf
 * @Date: 2020-08-19 16:46:49
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-20 16:26:25
 * @Description: plase write Description
 */
package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tengfei31/website/pkg/setting"
)


var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

//初始化db
func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "初始化database配置失败:%v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("DB_NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

//关闭db连接
func CloseDB() {
	defer db.Close()
}



