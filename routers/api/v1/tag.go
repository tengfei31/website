package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tengfei31/website/models"
	"github.com/tengfei31/website/pkg/e"
	"github.com/tengfei31/website/pkg/setting"
	"github.com/tengfei31/website/pkg/util"
	"github.com/unknwon/com"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
    name := c.Query("name")

    maps := make(map[string]interface{})
    data := make(map[string]interface{})

    if name != "" {
        maps["name"] = name
    }

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        maps["state"] = state
    }

    code := e.SUCCESS

    data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}

//新增文章标签
func AddTag(c *gin.Context) {}

//删除文章标签
func DelTag(c *gin.Context) {}

//修改文章标签
func EditTag(c *gin.Context) {}


