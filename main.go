/*
 * @Author: wtf
 * @Date: 2020-08-18 19:10:20
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-18 19:58:24
 * @Description: plase write Description
 */
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine{
	r := gin.Default()

	var data = make(map[string]string, 1)
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
	return r
}

func main() {
	r := setupRouter()
	r.Run("0.0.0.0:8088")
}
