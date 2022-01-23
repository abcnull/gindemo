package main

import (
	"gindemo/biz/handler"
	"gindemo/biz/handler/aboutme"
	"gindemo/biz/handler/account"
	"gindemo/biz/handler/center"
	"gindemo/biz/handler/person"
	"gindemo/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 创建一个 gin 实例
	r := gin.Default()

	/*--------- 静态资源加载 --------*/
	// 静态资源
	r.Static("/static", "static")

	/*--------- 加载前端模版 --------*/
	// 全部加载 view 下模版资源
	r.LoadHTMLGlob("view/**/*")

	// 鉴权中间件
	r.Use(middleware.JWTAuth())

	/*--------- 【首页模块】 --------*/
	// 首页 /
	r.GET("/", handler.HomeGet)

	/*--------- 【账户模块】 --------*/
	accountGroup := r.Group("/account")
	{
		/*--------- 注册 --------*/
		// 注册一个 get 请求 /register
		accountGroup.GET("/register", account.RegisterGet)
		// 注册一个 post 请求 /register
		accountGroup.POST("/register", account.RegisterPost)

		/*--------- 登录 --------*/
		// 登录一个 get 请求 /login
		accountGroup.GET("/login", account.LoginGet)
		// 登录一个 post 请求 /login
		accountGroup.POST("/login", account.LoginPost)

		/*--------- 退出登录 --------*/
		// 退出登录 get 请求 /logout
		accountGroup.GET("/logout", account.LogoutGet)
	}

	/*--------- 【个人中心模块】 --------*/
	// JWT token 鉴权中间件
	centerGroup := r.Group("/center")
	{
		centerGroup.GET("/", center.CenterGet)
	}

	/*--------- 【个人操作模块】 --------*/
	// JWT token 鉴权中间件
	personGroup := r.Group("/person")
	{
		// 文章列表页 OK
		personGroup.GET("/article", person.ArticleGet)
		// 具体文章页面 get 请求 /article/:id
		personGroup.GET("/article/:id", person.SpecificArticleGet)
		// 写文章页面 get 请求 /article/add
		personGroup.GET("/article/add", person.ArticleAddGet)
		// 写文章 post 请求 /article/add
		personGroup.POST("/article/add", person.ArticleAddPost)
		// 编辑文章页面 get 请求 /article/update/:id
		personGroup.GET("/article/update/:id", person.ArticleUpdateGet)
		// 编辑文章 post 请求 /article/update
		personGroup.POST("/article/update/:id", person.ArticleUpdatePost)
		// 删除文章 get 请求 /article/delete/:id
		personGroup.GET("/article/delete/:id", person.ArticleDeleteGet)

	}

	/*--------- 关于我模块 --------*/
	aboutmeGroup := r.Group("/aboutme")
	{
		aboutmeGroup.GET("/", aboutme.AboutMeGet)
	}

	return r
}
