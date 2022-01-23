package aboutme

import (
	"gindemo/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// AboutMeGet 访问关于我页面
func AboutMeGet(c *gin.Context) {
	// 拿到该用户 uid
	uid, isExist := c.Get("userId")
	if !isExist {
		c.Redirect(http.StatusMovedPermanently, "/account/login")
	}

	// 获取用户名
	name, err := service.QueryNameByUid(uid.(uint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "搜索用户出错",
		})
		return
	}

	// 获取用户文章数量
	artCount, err2 := service.QueryArticleCountByUid(uid.(uint64))
	if err2 != gorm.ErrRecordNotFound && err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "查找文章出错" + err.Error(),
		})
		return
	}

	// 返回 html 页面
	c.HTML(http.StatusOK, "aboutme.html", gin.H{
		"username": name,
		"artCount": artCount,
		"isLogin":  uid != nil,
	})
}
