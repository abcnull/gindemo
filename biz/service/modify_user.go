package service

import (
	"gindemo/biz/dal"
	"gindemo/biz/dal/entity"
)

// QueryUidByName 依据用户名拿到 uid
func QueryUidByName(name string) (uint64, error) {
	var user entity.User
	err := dal.DB.Where("username = ?", name).Take(&user).Error
	return uint64(user.ID), err
}

// QueryNameByUid 依据 uid 拿到用户名
func QueryNameByUid(uid uint64) (string, error) {
	var user entity.User
	err := dal.DB.Where("id = ?", uid).Take(&user).Error
	return user.Username, err
}

// QueryUserByName 依据用户名拿到用户
func QueryUserByName(name string) (entity.User, error) {
	var user entity.User
	err := dal.DB.Where("username = ?", name).Take(&user).Error
	return user, err
}

// AddNewUser 新增用户
func AddNewUser(user entity.User) error {
	return dal.DB.Create(&user).Error
}
