package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AboutMeGet 访问关于我页面
func AboutMeGet(c *gin.Context) {
	isLogin := IsSessionExist(c)

	if isLogin {
		c.HTML(http.StatusOK, "aboutme.html", gin.H{
			"isLogin":  isLogin,
			"username": GetSession(c, "loginUser"),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "你还没有登录！",
		})
	}
}
