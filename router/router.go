package router

import (
	"gindemo/controller"
	"gindemo/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	// 创建一个 gin 实例
	r := gin.Default()

	/*--------- 静态资源加载 --------*/
	// 静态资源
	r.Static("/static", "./static")

	/*--------- 加载前端模版 --------*/
	// 全部加载 view 下模版资源
	r.LoadHTMLGlob("view/**/*")

	/*--------- 中间件 --------*/
	// 给路由配置 session
	r.Use(middleware.UseSession())

	/*--------- 首页 --------*/
	// 首页 /
	r.GET("/", controller.HomeGet)

	/*--------- 注册 --------*/
	// 注册一个 get 请求 /register
	r.GET("/register", controller.RegisterGet)
	// 注册一个 post 请求 /register
	r.POST("/register", controller.RegisterPost)

	/*--------- 登录 --------*/
	// 登录一个 get 请求 /login
	r.GET("/login", controller.LoginGet)
	// 登录一个 post 请求 /login
	r.POST("/login", controller.LoginPost)
	// 退出登录 get 请求 /logout
	r.GET("/logout", controller.LogoutGet)

	/*--------- 文章 --------*/
	articleGroup := r.Group("/article")
	{
		// 文章页面 get 请求 /article/
		articleGroup.GET("/", controller.ArticleGet)
		// 具体文章页面 get 请求 /article/:id
		articleGroup.GET("/:id", controller.SpecificArticleGet)
		// 写文章页面 get 请求 /article/add
		articleGroup.GET("/add", controller.ArticleAddGet)
		// 写文章 post 请求 /article/add
		articleGroup.POST("/add", controller.ArticleAddPost)
		// 编辑文章页面 get 请求 /article/update/:id
		articleGroup.GET("/update/:id", controller.ArticleUpdateGet)
		// 编辑文章 post 请求 /article/update
		articleGroup.POST("update/:id", controller.ArticleUpdatePost)
		// 删除文章 get 请求 /article/delete/:id
		articleGroup.GET("/delete/:id", controller.ArticleDeleteGet)
	}

	/*--------- 关于我 --------*/
	r.GET("/aboutme", controller.AboutMeGet)

	return r
}
