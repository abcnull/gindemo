package account

import (
	"gindemo/biz/dal/entity"
	"gindemo/biz/model"
	"gindemo/biz/service"
	"gindemo/biz/service/check"
	"gindemo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// RegisterGet 注册页面
func RegisterGet(c *gin.Context) {
	uid, _ := c.Get("userId")
	c.HTML(http.StatusOK, "register.html", gin.H{
		"code":    0,
		"msg":     "注册页",
		"isLogin": uid != nil,
	})
}

// RegisterPost 处理注册
func RegisterPost(c *gin.Context) {
	// 接收表单数据
	req := new(model.RegisterPostReq)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "入参错误",
		})
		return
	}

	// 校验注册的输入是否正确
	isPass, resp := check.CheckRegister(req.Username, req.Password, req.RePassword)
	if !isPass {
		c.JSON(http.StatusOK, resp)
		return
	}

	// 插入一条用户数据
	user := entity.User{
		Username: req.Username,
		Phone:    "",
		Password: util.MD5(req.Password),
		Sex:      0,
		Birthday: time.Time{},
		Status:   0,
	}
	err := service.AddNewUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "注册失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "注册成功！请登录！",
		})
		c.Redirect(http.StatusMovedPermanently, "/account/login")
	}
}
