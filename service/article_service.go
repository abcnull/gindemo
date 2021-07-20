package service

import (
	"bytes"
	"gindemo/config"
	"gindemo/database"
	"gindemo/model"
	"html/template"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// MakeArticleListTemplate 依据页码创建文章列表模版
func MakeArticleListTemplate(isLogin bool, pageNum int) template.HTML {
	// 查询数据库获取文章列表
	articles, _ := model.QueryArticleWithPage(database.DB, pageNum)
	articlePage := ""
	// 遍历文章切片
	for _, article := range articles {
		// 修改 article content 内容
		if utf8.RuneCountInString(article.Content) > config.CHAR_PERARTICLE {
			article.Content = string([]rune(article.Content)[0:100]) + "..."
		}
		// 解析文件
		t, _ := template.ParseFiles("view/article/article_list.html")
		b := bytes.Buffer{}
		// 插入变量
		t.Execute(&b, map[string]interface{}{
			"Id":             article.Id,
			"Title":          article.Title,
			"Tags":           article.Tags,
			"Short":          article.Short,
			"Content":        article.Content,
			"Author":         article.Author,
			"CreateTime":     time.Unix(article.CreateTime, 0).Format("2006-01-02 15:04:05"),
			"LastUpdateTime": time.Unix(article.LastUpdateTime, 0).Format("2006-01-02 15:04:05"),
			"isLogin":        isLogin,
		})
		//t.Execute(&b, isLogin)
		articlePage += b.String()
	}
	return template.HTML(articlePage)
}

// MakePageBarArgs 创建页码参数
func MakePageBarArgs(pageNum int) map[string]interface{} {
	// 拿到文章总数
	allNum := model.QueryArticleCount(database.DB)
	if pageNum <= 0 {
		pageNum = 1
	}
	pageBarMap := make(map[string]interface{})

	pageBarMap["pageNum"] = pageNum
	// 判断首页
	if pageNum == 1 {
		pageBarMap["isFirstPage"] = true
		pageBarMap["prePageNum"] = strconv.Itoa(pageNum)
	} else {
		pageBarMap["isFirstPage"] = false
		pageBarMap["prePageNum"] = strconv.Itoa(pageNum - 1)
	}
	// 判断尾页
	if pageNum*config.SIZE_PERPAGE >= allNum {
		pageBarMap["isLastPage"] = true
		pageBarMap["nextPageNum"] = strconv.Itoa(pageNum)
	} else {
		pageBarMap["isLastPage"] = false
		pageBarMap["nextPageNum"] = strconv.Itoa(pageNum + 1)
	}
	return pageBarMap
}

// MakeSpecificArticleArgs 获取单篇文章界面中的参数信息
func MakeSpecificArticleArgs(id int) (model.Article, string, string) {
	// 根据 id 查数据库
	article := model.QueryArticleWithId(database.DB, id)
	createTime := time.Unix(article.CreateTime, 0).Format("2006-01-02 15:04:05")
	lastUpdateTime := time.Unix(article.LastUpdateTime, 0).Format("2006-01-02 15:04:05")
	return article, createTime, lastUpdateTime
}

// ChangeNToBr 转换 content 为 <div> 可识别的换行符号
func ChangeNToBr(article model.Article) model.Article {
	// 将文本内容替换成可以显示换行的文本
	article.Content = strings.Replace(article.Content, "\r\n", "<br>", -1)
	article.Content = strings.Replace(article.Content, "\n", "<br>", -1)
	return article
}

// ChangeNToNextLine 转换 content 为 text 中可识别换行符
func ChangeNToNextLine(article model.Article) model.Article {
	// 格式化 content
	article.Content = strings.Replace(article.Content, "\r\n", "&#13;&#10;", -1)
	article.Content = strings.Replace(article.Content, "\n", "&#13;&#10;", -1)
	return article
}

// UpdateArticleProcess 更新文章
func UpdateArticleProcess(id int, args ...interface{}) (int64, error) {
	// 依据 id 获取 article
	article := model.QueryArticleWithId(database.DB, id)

	// article 重新赋值
	article.Title = args[0].(string)
	article.Tags = args[1].(string)
	article.Short = args[2].(string)
	article.Content = args[3].(string)
	article.LastUpdateTime = args[4].(int64)

	// 更新数据
	return model.UpdateArticle(database.DB, article)
}

// AddArticleProcess 新增一片文章
func AddArticleProcess(args ...interface{}) (int64, error) {
	// 实例化 Article
	article := model.Article{Title: args[0].(string), Tags: args[1].(string), Short: args[2].(string), Content: args[3].(string), Author: args[4].(string),
		CreateTime: args[5].(int64), LastUpdateTime: args[6].(int64)}
	return model.InsertArticle(database.DB, article)
}
