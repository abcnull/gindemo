package check

import (
	"gindemo/biz/service"
	"gindemo/util"
)

// CheckNameFormat 校验用户名格式
func CheckNameFormat(name string) bool {
	// 校验用户名不为空
	return name != ""
}

// CheckPwdFormat 校验密码格式
func CheckPwdFormat(pwd string) bool {
	// 校验密码 6-12 位
	return len(pwd) >= 6 && len(pwd) <= 12
}

// ComparePwdAndRePwd 验证原密码和重新输出密码是否匹配
func ComparePwdAndRePwd(pwd string, rePwd string) bool {
	// 校验原密码和重新输入密码相等
	return pwd == rePwd
}

// IsUserExist 验证用户是否在数据库存在
func IsUserExist(name string) (bool, error) {
	// 校验用户名不存在
	if _, err := service.QueryUidByName(name); err != nil {
		return false, err
	} else {
		return true, err
	}
}

// IsUserAndPwdRight 验证用户密码在数据库是否存在且匹配
func IsUserAndPwdRight(name string, pwd string) (bool, error) {
	// 查询到用户名
	if user, err := service.QueryUserByName(name); err != nil {
		return false, err
	} else {
		// 比较密码
		if util.MD5(pwd) == user.Password {
			return true, err
		} else {
			return false, err
		}
	}
}
