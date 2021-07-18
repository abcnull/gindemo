package service

import (
	"gindemo/database"
	"gindemo/model"
	"gindemo/util"
	"time"
)

// JudgeUserExist 判定用户是否存在

// AddNewUserProcess 新增用户
func AddNewUserProcess(args ...interface{}) (int64, error) {
	// 用户数据
	user := model.User{Username: args[0].(string), Password: util.MD5(args[1].(string)), CreateTime: time.Now().Unix()}

	// 返回
	return model.InsertUser(database.DB, user)
}
