package controller

import (
	"gindemo/database"
	"gindemo/model"
	"gindemo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterGet 注册页面
func RegisterGet(c *gin.Context) {
	// 拿到全部文章数量
	artNum := model.QueryArticleCount(database.DB)

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title":  "注册页",
		"artNum": artNum,
	})
}

// RegisterPost 处理注册
func RegisterPost(c *gin.Context) {
	// 获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	rePassword := c.PostForm("rePassword")

	// 返回结果
	resp := gin.H{}

	// 重复输入密码错误
	if password != rePassword {
		resp = gin.H{
			"code":    1,
			"message": "两次密码不一致",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// 注册之前先判断用户是否已经被注册，若被注册返回错误
	flag := service.JudgeUserExist(username, password)
	if flag {
		resp = gin.H{
			"code":    1,
			"message": "用户名已经存在",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// 插入一条用户数据
	_, err := service.AddNewUserProcess(username, password)

	if err != nil {
		resp = gin.H{
			"code":    1,
			"message": "注册失败",
		}
	} else {
		resp = gin.H{
			"code":    0,
			"message": "注册成功！请登录",
		}
	}
	c.JSON(http.StatusOK, resp)
}
