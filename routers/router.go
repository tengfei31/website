/*
 * @Author: wtf
 * @Date: 2020-08-19 20:15:28
 * @LastEditors: wtf
 * @LastEditTime: 2020-08-25 12:05:05
 * @Description: plase write Description
 */
package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tengfei31/website/middleware/jwt"
	"github.com/tengfei31/website/pkg/setting"
	"github.com/tengfei31/website/routers/api"
	v1 "github.com/tengfei31/website/routers/api/v1"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(setting.RunMode)
	

	//授权路由
	r.GET("auth", api.GetAuth)

	//注册路由
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
    {
		//标签路由
        //获取标签列表
        apiv1.GET("/tags", v1.GetTags)
        //新建标签
        apiv1.POST("/tags", v1.AddTag)
        //更新指定标签
        apiv1.PUT("/tags/:id", v1.EditTag)
        //删除指定标签
		apiv1.DELETE("/tags/:id", v1.DelTag)
		
		//文章路由
		//获取文章列表
        apiv1.GET("articles", v1.GetArticles)
        //获取指定文章
        apiv1.GET("articles/:id", v1.GetArticle)
        //新建文章
        apiv1.POST("articles", v1.AddArticle)
        //更新指定文章
        apiv1.PUT("articles/:id", v1.EditArticle)
        //删除指定文章
        apiv1.DELETE("articles/:id", v1.DeleteArticle)
	}
	
	
	r.GET("ping", func(c *gin.Context) {
		var data = make(map[string]string)
		data["name"] = "wtf"
		data["age"] = "男"
		data["sex"] = "27"
		//c.String(http.StatusOK, "pong")
		c.JSON(http.StatusOK, gin.H{
			"msg":"success",
			"code":0,
			"data":data,
		})
	})
	
	return r
}

