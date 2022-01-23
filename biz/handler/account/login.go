package account

import (
	"fmt"
	"gindemo/biz/model"
	"gindemo/biz/service"
	"gindemo/biz/service/check"
	"gindemo/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// LoginGet 登录页面
func LoginGet(c *gin.Context) {
	// uid
	uid, _ := c.Get("userId")

	// 返回 html
	resp := gin.H{
		"code":    0,
		"msg":     "登录页",
		"isLogin": uid != nil,
	}
	c.HTML(http.StatusOK, "login.html", resp)
}

// LoginPost 登录请求
func LoginPost(c *gin.Context) {
	req := new(model.LoginPostReq)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "入参错误",
		})
		return
	}

	// 校验登录
	isPass, resp := check.CheckLogin(req.Username, req.Password)
	if !isPass {
		c.JSON(http.StatusOK, resp)
		return
	}

	// 生成 token 并做响应头返回
	uid, err := service.QueryUidByName(req.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "数据库查询错误" + err.Error(),
		})
		return
	}
	expireTime := time.Now().Add(60 * 24 * time.Hour)
	tokenStr, err := middleware.GenerateToken(uid, expireTime)
	c.SetCookie("Authorization", tokenStr, 3600, "/", "127.0.0.1:8081", false, true)
	c.Set("userId", uid)
	fmt.Println("这是loginpost")
	fmt.Println(c.Get("userId"))
	c.JSON(http.StatusOK, resp)
	c.Redirect(http.StatusMovedPermanently, "/")
}
