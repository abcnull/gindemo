package handler

import (
	"gindemo/biz/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HomeGet 首页
func HomeGet(c *gin.Context) {
	var name string
	var err error
	uid, isExist := c.Get("userId")
	if isExist {
		name, err = service.QueryNameByUid(uid.(uint64))
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "/")
		}
	}

	// 返回 html
	resp := gin.H{
		"code":     0,
		"msg":      "首页",
		"isLogin":  uid != nil,
		"username": name,
	}
	c.HTML(http.StatusOK, "home.html", resp)
}
