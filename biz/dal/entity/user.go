package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User 用户
type User struct {
	gorm.Model
	Username string    `json:"username" form:"username"` // 用户名
	Phone    string    // 手机号
	Password string    `json:"password" form:"password"` // 加密后的密码
	Sex      int8      // 性别
	Birthday time.Time // 生日
	Status   int8      `json:"status" form:"status"` // 状态
}
