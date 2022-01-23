package person

import (
	"gindemo/biz/dal/entity"
	"gindemo/biz/model"
	"gindemo/biz/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"html/template"
	"net/http"
	"strconv"
)

/*------------------------ 文章列表页面 -------------------------*/

// ArticleGet 根据 query 参数 page 获取文章页面 Get
func ArticleGet(c *gin.Context) {
	// 拿到该用户 uid
	uid, isExist := c.Get("userId")
	if !isExist {
		// 如果用户不存在重定向登录
		c.Redirect(http.StatusMovedPermanently, "/account/login")
		return
	}

	// 拿到 pageNum 页数
	req := new(model.ArticleGetReq)
	err := c.ShouldBindQuery(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "错误 page",
		})
		return
	}

	// 拿到该用户 uid 的所有文章
	artCount, err := service.QueryArticleCountByUid(uid.(uint64))
	if err != gorm.ErrRecordNotFound && err != nil {
		// 如果查询单纯报错
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "用 uid 查询文章数据库报错",
		})
		return
	}

	// 生成该用户的文章列表模版 template.HTML
	listContent, err := service.MakeUserArticleListTemplate(uid.(uint64), req.PageNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "展示前端用户文章列表出现错误",
		})
		return
	}
	// 页码参数 map 生成
	pageBarMap, err := service.MakePageBarArgs(uid.(uint64), req.PageNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "页码数字生成有问题",
		})
	}

	// 依据 uid 查询 username
	name, err := service.QueryNameByUid(uid.(uint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "查询用户名出错",
		})
		return
	}

	// 返回
	c.HTML(http.StatusOK, "article.html", gin.H{
		// 用户名
		"username": name,
		// 文章数量
		"artCount": artCount,
		// 文章列表模版
		"listContent": listContent,
		// 页码
		"isFirstPage": pageBarMap["isFirstPage"],
		"isLastPage":  pageBarMap["isLastPage"],
		"pageNum":     pageBarMap["pageNum"],
		"prePageNum":  pageBarMap["prePageNum"],
		"nextPageNum": pageBarMap["nextPageNum"],
		// 登录与否
		"isLogin": true,
	})
}

// ArticleDeleteGet 删除某个文章 Get
func ArticleDeleteGet(c *gin.Context) {
	// 获取路由参数表明是哪一篇文章
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "文章不存在",
		})
	}

	// 删除文章
	if err := service.DeleteArticleById(uint64(id)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "删除文章失败",
		})
		return
	}

	// 返回
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "删除文章成功",
	})
	c.Redirect(http.StatusMovedPermanently, "/person/article")
}

/*------------------------ 单篇文章阅读页面 -------------------------*/

// SpecificArticleGet 根据路由中参数，获取具体文章页面 Get
func SpecificArticleGet(c *gin.Context) {
	// 获取路由参数表明是哪一篇文章 id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 {
		// id 无法转化成正常数字
		c.JSON(http.StatusNotFound, gin.H{
			"code": 1,
			"msg":  "入参有问题",
		})
		return
	}

	// 拿到 uid
	uid, ifExist := c.Get("userId")
	if !ifExist {
		c.Redirect(http.StatusMovedPermanently, "/account/login")
		return
	}
	username, err := service.QueryNameByUid(uid.(uint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "查询用户出现错误",
		})
		return
	}

	// 拿到该用户全部文章数量
	artCount, err := service.QueryArticleCountByUid(uid.(uint64))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "查询用户文章出现错误",
		})
		return
	}

	// 拿到 id 对应的 article
	article, err := service.QueryArticleById(uint64(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "查询用户文章出现错误",
		})
		return
	}
	// format tag
	article.Tags = service.ChangeTags(article.Tags)
	// format content, 把 \n 换成 <br>
	article = service.ChangeNToBr(article)

	// 渲染 html
	resp := gin.H{
		"username":      username,
		"artCount":      artCount,
		"article":       article,
		"contentFormat": template.HTML(article.Content),
		"tags":          template.HTML(article.Tags),
		"isLogin":       true,
	}
	c.HTML(http.StatusOK, "specific_article.html", resp)
}

/*------------------------ 单篇文章编辑页面 -------------------------*/

// ArticleUpdateGet 进入编辑文章页面 Get
func ArticleUpdateGet(c *gin.Context) {
	// 拿到 uid
	uid, isExist := c.Get("userId")
	if !isExist {
		c.Redirect(http.StatusMovedPermanently, "/account/login")
		return
	}
	name, err := service.QueryNameByUid(uid.(uint64))
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/account/login")
		return
	}

	// 获取路由参数表明是哪一篇文章
	idStr := c.Param("id")

	// 拿到用户全部文章数量
	artCount, err := service.QueryArticleCountByUid(uid.(uint64))
	if err != gorm.ErrRecordNotFound && err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "搜不到用户文章",
		})
		return
	}

	// 依据 id 搜索 article
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "参数有问题",
		})
		return
	}
	article, err := service.QueryArticleById(uint64(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "搜索文章出错",
		})
		return
	}
	// 把 \n 换成 &#13;&#10;
	article = service.ChangeNToNextLine(article)

	// 返回
	c.HTML(http.StatusOK, "write.html", gin.H{
		"username":      name,
		"artCount":      artCount,
		"isWrite":       true,
		"article":       article,
		"contentFormat": template.HTML(article.Content),
		"isLogin":       true,
	})
}

// ArticleUpdatePost 编辑文章提交 Post
func ArticleUpdatePost(c *gin.Context) {
	// 取数据
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "没有此文章",
		})
	}

	req := new(model.ArticleChangeReq)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "入参错误",
		})
		return
	}

	// article
	article := entity.Article{
		Title:   req.Title,
		Tags:    req.Tags,
		Short:   req.Short,
		Content: req.Content,
	}

	// 依据 id 更新文章
	if err := service.UpdateArticleById(article, uint64(id)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "更新文章出错",
		})
		return
	}

	// 返回
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "更新文章成功",
	})
}

/*------------------------ 写文章页面 -------------------------*/

// ArticleAddGet 写文章的页面 Get
func ArticleAddGet(c *gin.Context) {
	// 拿到 uid
	uid, isExist := c.Get("userId")
	if !isExist {
		c.Redirect(http.StatusMovedPermanently, "/account/login")
	}
	username, err := service.QueryNameByUid(uid.(uint64))
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/account/login")
	}

	// 拿到该用户全部文章数量
	artCount, err := service.QueryArticleCountByUid(uid.(uint64))
	if err != gorm.ErrRecordNotFound && err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"msg":     "搜索用户文章出错",
			"isLogin": true,
		})
	}

	// 返回模版
	c.HTML(http.StatusOK, "write.html", gin.H{
		"username": username,
		"artCount": artCount,
		"isLogin":  true,
	})
}

// ArticleAddPost 提交文章请求 Post
func ArticleAddPost(c *gin.Context) {
	// 获取 uid
	uid, isExist := c.Get("userId")
	if !isExist {
		c.Redirect(http.StatusMovedPermanently, "/account/login")
		return
	}

	req := new(model.ArticleChangeReq)
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "入参错误",
		})
		return
	}

	article := entity.Article{
		Uid:     uid.(uint64),
		Title:   req.Title,
		Tags:    req.Tags,
		Short:   req.Short,
		Content: req.Content,
	}

	// 更新数据
	if err := service.AddArticle(article); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "添加文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "添加文章成功",
	})
}
