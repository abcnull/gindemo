package main

import (
	"gindemo/database"
	"gindemo/router"
)

func main() {
	// 初始化数据表
	database.InitMysql()

	// 创建一个路由实例
	r := router.InitRouter()

	// 运行服务
	r.Run(":8081")
}
