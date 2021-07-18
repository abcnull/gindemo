package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HomeGet 首页
func HomeGet(c *gin.Context) {
	// session 是否存在
	isLogin := IsSessionExist(c)
	// 返回 html
	resp := gin.H{
		"isLogin":  isLogin,
		"username": GetSession(c, "loginUser"),
	}
	c.HTML(http.StatusOK, "home.html", resp)
}
