/*
 * @Author: wtf
 * @Date: 2020-08-19 20:15:28
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-19 20:31:12
 * @Description: plase write Description
 */
package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tengfei31/website/pkg/setting"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(setting.RunMode)
	
	//注册路由
	var data = make(map[string]string)
	data["name"] = "wtf"
	data["age"] = "男"
	data["sex"] = "27"
	r.GET("ping", func(c *gin.Context) {
		//c.String(http.StatusOK, "pong")
		c.JSON(http.StatusOK, gin.H{
			"msg":"success",
			"code":0,
			"data":data,
		})
	})

	r.GET("test", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"message" : "test",
		})
	})
	
	return r
}

