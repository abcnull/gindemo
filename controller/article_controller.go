package controller

import (
	"gindemo/database"
	"gindemo/model"
	"gindemo/service"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

/*------------------------ 文章列表页面 -------------------------*/

// ArticleGet 根据 query 参数 page 获取文章页面 Get
func ArticleGet(c *gin.Context) {
	// 拿到 pageNum 页数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// 判断是否登录
	isLogin := IsSessionExist(c)
	// 拿到全部文章数量
	artNum := model.QueryArticleCount(database.DB)

	// 文章列表模版
	listContent := service.MakeArticleListTemplate(isLogin, pageNum)
	// 页码参数 map 生成
	pageBarMap := service.MakePageBarArgs(pageNum)

	// 返回
	resp := gin.H{
		// 是否登录
		"isLogin":  isLogin,
		"username": GetSession(c, "loginUser"),
		// 文章数量
		"artNum": artNum,
		// 文章列表
		"listContent": listContent,
		// 页码
		"isFirstPage": pageBarMap["isFirstPage"],
		"isLastPage":  pageBarMap["isLastPage"],
		"pageNum":     pageBarMap["pageNum"],
		"prePageNum":  pageBarMap["prePageNum"],
		"nextPageNum": pageBarMap["nextPageNum"],
	}
	c.HTML(http.StatusOK, "article.html", resp)
}

// ArticleDeleteGet 删除某个文章 Get
func ArticleDeleteGet(c *gin.Context) {
	// 获取路由参数表明是哪一篇文章
	idStr := c.Param("id")

	// 删除文章
	if id, err := strconv.Atoi(idStr); err == nil {
		model.DeleteArticleWithId(database.DB, id)
	} else {
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "/article")
		return
	}

	// 返回
	c.Redirect(http.StatusMovedPermanently, "/article")
}

/*------------------------ 单篇文章阅读页面 -------------------------*/

// SpecificArticleGet 根据路由中参数，获取具体文章页面 Get
func SpecificArticleGet(c *gin.Context) {
	// 获取路由参数表明是哪一篇文章
	idStr := c.Param("id")
	// 判断登录
	isLogin := IsSessionExist(c)
	// 拿到全部文章数量
	artNum := model.QueryArticleCount(database.DB)

	// 转换 id 获取关键 args
	article := model.Article{}
	createTime := ""
	lastUpdateTime := ""
	if id, err := strconv.Atoi(idStr); err == nil && id > 0 {
		article, createTime, lastUpdateTime = service.MakeSpecificArticleArgs(id)
	} else {
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "/article")
		return
	}

	// 把 \n 换成 <br>
	article = service.ChangeNToBr(article)

	// 渲染 html
	resp := gin.H{
		"isLogin":        isLogin,
		"username":       GetSession(c, "loginUser"),
		"artNum":         artNum,
		"article":        article,
		"createTime":     createTime,
		"lastUpdateTime": lastUpdateTime,
		"contentFormat":  template.HTML(article.Content),
		"tags":           template.HTML(article.Tags),
	}
	c.HTML(http.StatusOK, "specific_article.html", resp)
}

/*------------------------ 单篇文章编辑页面 -------------------------*/

// ArticleUpdateGet 进入编辑文章页面 Get
func ArticleUpdateGet(c *gin.Context) {
	// 获取路由参数表明是哪一篇文章
	idStr := c.Param("id")
	// 判断登录
	isLogin := IsSessionExist(c)
	// 拿到全部文章数量
	artNum := model.QueryArticleCount(database.DB)

	article := model.Article{}
	if id, err := strconv.Atoi(idStr); err == nil {
		article = model.QueryArticleWithId(database.DB, id)
	} else {
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "/article")
		return
	}

	// 把 \n 换成 &#13;&#10;
	article = service.ChangeNToNextLine(article)

	// 返回
	resp := gin.H{
		"isLogin":       isLogin,
		"username":      GetSession(c, "loginUser"),
		"artNum":        artNum,
		"isWrite":       true,
		"article":       article,
		"contentFormat": template.HTML(article.Content),
	}
	c.HTML(http.StatusOK, "write.html", resp)
}

// ArticleUpdatePost 编辑文章提交 Post
func ArticleUpdatePost(c *gin.Context) {
	// 取数据
	idStr := c.Param("id")
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	lastUpdateTime := time.Now().Unix()

	resp := gin.H{}
	if id, err := strconv.Atoi(idStr); err == nil {
		_, err2 := service.UpdateArticleProcess(id, title, tags, short, content, lastUpdateTime)
		// 返回结果
		if err2 != nil {
			resp = gin.H{
				"code":    1,
				"message": "更新失败！",
			}
		} else {
			resp = gin.H{
				"code":    0,
				"message": "更新成功",
			}
		}
	} else {
		// 重定向
		c.Redirect(http.StatusMovedPermanently, "/article")
		return
	}

	// 返回
	c.JSON(http.StatusOK, resp)
}

/*------------------------ 写文章页面 -------------------------*/

// ArticleAddGet 写文章的页面 Get
func ArticleAddGet(c *gin.Context) {
	// 登录信息
	isLogin := IsSessionExist(c)
	// 拿到全部文章数量
	artNum := model.QueryArticleCount(database.DB)
	resp := gin.H{
		"isLogin":  isLogin,
		"username": GetSession(c, "loginUser"),
		"artNum":   artNum,
	}
	c.HTML(http.StatusOK, "write.html", resp)
}

// ArticleAddPost 提交文章请求 Post
func ArticleAddPost(c *gin.Context) {
	// 获取表单信息
	title := c.PostForm("title")
	author := GetSession(c, "loginUser").(string)
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	createTime := time.Now().Unix()
	lastUpdateTime := createTime

	// 更新数据
	_, err := service.AddArticleProcess(title, author, tags, short, content, createTime, lastUpdateTime)

	// 返回数据
	resp := gin.H{}
	if err != nil {
		resp = gin.H{
			"code":    1,
			"message": "新增文章异常！",
		}
	} else {
		resp = gin.H{
			"code":    0,
			"message": "成功新增一篇文章",
		}
	}
	c.JSON(http.StatusOK, resp)
}
