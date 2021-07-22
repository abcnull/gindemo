package controller

import (
	"gindemo/database"
	"gindemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AboutMeGet 访问关于我页面
func AboutMeGet(c *gin.Context) {
	// 登录状态
	isLogin := IsSessionExist(c)
	// 拿到全部文章数量
	artNum := model.QueryArticleCount(database.DB)

	if isLogin {
		c.HTML(http.StatusOK, "aboutme.html", gin.H{
			"isLogin":  isLogin,
			"username": GetSession(c, "loginUser"),
			"artNum":   artNum,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "你还没有登录！",
		})
	}
}
