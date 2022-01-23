package check

import (
	"github.com/gin-gonic/gin"
)

// CheckLogin 校验登录
func CheckLogin(name string, pwd string) (bool, gin.H) {
	// 验证用户名不为空
	if !CheckNameFormat(name) {
		return false, gin.H{
			"code": 1,
			"msg":  "用户名格式错误",
		}
	}
	// 验证密码位数 [6, 12]
	if !CheckPwdFormat(pwd) {
		return false, gin.H{
			"code": 1,
			"msg":  "密码格式错误",
		}
	}
	// 验证用户密码在数据库中是否匹配存在
	if is, _ := IsUserAndPwdRight(name, pwd); !is {
		return false, gin.H{
			"code": 1,
			"msg":  "用户名或密码错误",
		}
	}
	return true, gin.H{
		"code": 0,
		"msg":  "登录成功",
	}
}
