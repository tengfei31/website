/*
 * @Author: wtf
 * @Date: 2020-08-24 19:22:33
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-25 12:11:32
 * @Description: plase write Description
 */
package jwt

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tengfei31/website/pkg/e"
	"github.com/tengfei31/website/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
