package service

import (
	"gindemo/database"
	"gindemo/model"
	"gindemo/util"
)

// JudgeUserExist 判定用户是否存在
func JudgeUserExist(username string, password string) bool {
	// md5 转换密码
	password = util.MD5(password)
	// 查到 id
	id := model.QueryIdWithUserPwd(database.DB, username, password)
	// 返回
	if id > 0 {
		return true
	} else {
		return false
	}
}
