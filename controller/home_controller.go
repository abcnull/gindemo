package controller

import (
	"gindemo/database"
	"gindemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HomeGet 首页
func HomeGet(c *gin.Context) {
	// session 是否存在
	isLogin := IsSessionExist(c)
	// 拿到全部文章数量
	artNum := model.QueryArticleCount(database.DB)

	// 返回 html
	resp := gin.H{
		"isLogin":  isLogin,
		"username": GetSession(c, "loginUser"),
		"artNum":   artNum,
	}
	c.HTML(http.StatusOK, "home.html", resp)
}
