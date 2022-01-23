package account

import (
	"gindemo/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// LogoutGet 退出登录
func LogoutGet(c *gin.Context) {
	// 更改密钥
	util.ChangeSecretKey()

	log.Println("+++", util.GetSecretKey())
	// 重定向
	c.Redirect(http.StatusMovedPermanently, "/")
}
