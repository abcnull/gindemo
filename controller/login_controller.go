package controller

import (
	"gindemo/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginGet 登录页面
func LoginGet(c *gin.Context) {
	// 返回 html
	resp := gin.H{
		"title": "登录页",
	}
	c.HTML(http.StatusOK, "login.html", resp)
}

// LoginPost 登录请求
func LoginPost(c *gin.Context) {
	// 表单接收数据
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 判定用户存在
	flag := service.JudgeUserExist(username, password)

	// 返回结果
	resp := gin.H{}
	if flag {
		// session
		s := sessions.Default(c)
		// session 设置
		s.Set("loginUser", username)
		// 无论是 Set 还是 Delete 都需要 Save 函数保存
		s.Save()
		resp = gin.H{
			"code":    0,
			"message": "登录成功，欢迎进入！",
		}
	} else {
		resp = gin.H{
			"code":    1,
			"message": "用户名或密码错误！",
		}
	}
	c.JSON(http.StatusOK, resp)
}

// LogoutGet 退出登录
func LogoutGet(c *gin.Context) {
	// 清除 session 信息
	s := sessions.Default(c)
	s.Delete("loginUser")
	s.Save()
	// 重定向
	c.Redirect(http.StatusMovedPermanently, "/")
}
