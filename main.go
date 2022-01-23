package main

import (
	"gindemo/biz/dal"
	"gindemo/config"
	"gindemo/util"
)

func main() {
	// 配置文件加载
	config.Init()

	// 初始化数据表
	dal.InitMysql()
	defer dal.DB.Close()

	// 初始化 secretKey
	util.InitSecretKey()

	// 创建一个路由实例，其中进行路由注册
	r := InitRouter()

	// 运行服务
	r.Run(":8081")
}
