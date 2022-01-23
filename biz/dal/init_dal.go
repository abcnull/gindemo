package dal

import (
	"fmt"
	"gindemo/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

// DB 数据库
var DB *gorm.DB

// InitMysql 初始化 mysql 数据库
func InitMysql() {
	driver := config.DB_DRIVER

	username := config.DB_USER
	password := config.DB_PWD
	host := config.DB_ADD
	port := config.DB_PORT
	database := config.DB_NAME
	charset := config.DB_CHARSET

	// 数据源名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	// 连接数据库
	var err error
	DB, err = gorm.Open(driver, dsn)
	if err != nil {
		panic("无法连接数据库！" + err.Error())
	}

	// gorm 默认识别表名在后头加 s，比如比你的相查 user 表，结果发现 gorm 搜的是 users 表，这里使其不加 s
	DB.SingularTable(true)

	// 设置连接池
	DB.DB().SetMaxIdleConns(20)
	// 设置最大连接数
	DB.DB().SetMaxOpenConns(100)
	// 设置连接时间
	DB.DB().SetConnMaxLifetime(time.Second * 30)
}
