package entity

import (
	"github.com/jinzhu/gorm"
)

// Article 文章
type Article struct {
	gorm.Model
	Title   string // 文章标题
	Tags    string // 文章类型
	Short   string // 文章简介
	Content string // 文章内容
	Uid     uint64 `gorm:"column:userid"` // 作者
}
