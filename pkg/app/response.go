/*
 * @Author: wtf
 * @Date: 2020-08-28 22:57:48
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-28 23:01:59
 * @Description: plase write Description
 */
package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tengfei31/website/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, errorCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code" :errorCode,
		"msg": e.GetMsg(errorCode),
		"data" : data,
	})
	return
}



