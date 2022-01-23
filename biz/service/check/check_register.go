package check

import (
	"github.com/gin-gonic/gin"
)

// CheckRegister 校验注册的用户密码是否符合要求
func CheckRegister(name string, pwd string, rePwd string) (bool, gin.H) {
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
	// 验证原密码和重复输入密码相等
	if !ComparePwdAndRePwd(pwd, rePwd) {
		return false, gin.H{
			"code": 1,
			"msg":  "两次输入密码不一致",
		}
	}
	// 验证用户在数据库中是否存在
	if is, _ := IsUserExist(name); !is {
		return false, gin.H{
			"code": 1,
			"msg":  "用户已存在",
		}
	}
	return true, gin.H{
		"code": 0,
		"msg":  "注册成功",
	}
}
