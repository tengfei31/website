/*
 * @Author: wtf
 * @Date: 2020-08-19 16:46:49
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-27 20:18:45
 * @Description: plase write Description
 */
package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tengfei31/website/pkg/setting"
)


var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn int `json:"deleted_on"`
}

//初始化db
func Setup() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = setting.DataBaseSetting.Type
	dbName = setting.DataBaseSetting.DbName
	user = setting.DataBaseSetting.User
	password = setting.DataBaseSetting.Password
	host = setting.DataBaseSetting.Host
	tablePrefix = setting.DataBaseSetting.TablePrefix
	var url string = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName) 
	db, err = gorm.Open(dbType, url)
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	    return tablePrefix + defaultTableName;
	}
	
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

//关闭db连接
func CloseDB() {
	defer db.Close()
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime:= time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	} 
}

func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		var sql string
		if !scope.Search.Unscoped && hasDeletedOnField {
			sql = fmt.Sprintf("UPDATE %v SET %v=%v%v%v", 
				scope.QuotedTableName(), 
				scope.Quote(deletedOnField.DBName), 
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption))
		} else {
			sql = fmt.Sprintf("DELETE FROM %v%v%v", 
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption))
		}
		scope.Raw(sql).Exec()
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}





