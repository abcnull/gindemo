package model

import (
	"database/sql"
	"gindemo/config"
	"gindemo/database"
	"strconv"
)

// Article 文章
type Article struct {
	Id             int
	Title          string
	Tags           string
	Short          string
	Content        string
	Author         string
	CreateTime     int64
	LastUpdateTime int64
}

// ----------------- 数据库操作 ----------------------

// ------ 查询 -------

// QueryArticleCount 查询文章总数
func QueryArticleCount(DB *sql.DB) int {
	sql := "select count(*) from article"
	num := 0
	DB.QueryRow(sql).Scan(&num)
	return num
}

// QueryArticleWithId 根据 id 查询文章
func QueryArticleWithId(DB *sql.DB, id int) Article {
	sql := "select * from article where id=" + strconv.Itoa(id)
	article := Article{}
	row := DB.QueryRow(sql)
	{
		id := 0
		title := ""
		author := ""
		tags := ""
		short := ""
		content := ""
		var createTime int64 = 0
		var lastUpdateTime int64 = 0
		row.Scan(&id, &title, &author, &tags, &short, &content, &createTime, &lastUpdateTime)
		article = Article{Id: id, Title: title, Author: author, Tags: tags, Short: short, Content: content,
			CreateTime: createTime, LastUpdateTime: lastUpdateTime}
	}
	return article
}

// QueryArticleWithPage 根据页码查询文章
func QueryArticleWithPage(DB *sql.DB, pageNum int) ([]Article, error) {
	if pageNum <= 0 {
		pageNum = 1
	}
	fromIndexStr := strconv.Itoa((pageNum - 1) * config.SIZE_PERPAGE)
	pageNumStr := strconv.Itoa(config.SIZE_PERPAGE)
	sql := "select * from article order by lastupdatetime desc limit " + fromIndexStr + ", " + pageNumStr
	// 分页查询文章
	rows, err := database.QueryRowsDB(DB, sql)
	var articles []Article
	for rows.Next() {
		id := 0
		title := ""
		author := ""
		tags := ""
		short := ""
		content := ""
		var createTime int64 = 0
		var lastUpdateTime int64 = 0
		rows.Scan(&id, &title, &author, &tags, &short, &content, &createTime, &lastUpdateTime)
		articles = append(articles, Article{Id: id, Title: title, Short: short, Content: content, Author: author,
			CreateTime: createTime, LastUpdateTime: lastUpdateTime})
	}
	return articles, err
}

// ------ 插入 -------

// InsertArticle 插入一片文章
func InsertArticle(DB *sql.DB, article Article) (int64, error) {
	sql := "insert into article(title,tags,short,content,author,createtime,lastupdatetime) values(?,?,?,?,?,?,?)"
	return database.ModifyDB(DB, sql, article.Title, article.Tags, article.Short, article.Content, article.Author,
		article.CreateTime, article.LastUpdateTime)
}

// ------ 更新 -------

// UpdateArticle 更新一片文章
func UpdateArticle(DB *sql.DB, article Article) (int64, error) {
	sql := "update article set title=?, tags=?, short=?, content=?, lastupdatetime=? where id=?"
	return database.ModifyDB(DB, sql, article.Title, article.Tags, article.Short, article.Content,
		article.LastUpdateTime, article.Id)
}

// ------ 删除 -------

// DeleteArticleWithId 依据 id 删除文章
func DeleteArticleWithId(DB *sql.DB, id int) (int64, error) {
	sql := "delete from article where id=" + strconv.Itoa(id)
	return database.ModifyDB(DB, sql)
}
