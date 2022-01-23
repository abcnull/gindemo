package center

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CenterGet 个人中心页面
func CenterGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "个人中心页",
	})
}
