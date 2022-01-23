package service

import (
	"bytes"
	"errors"
	"gindemo/biz/dal"
	"gindemo/biz/dal/entity"
	"gindemo/config"
	"github.com/jinzhu/gorm"
	"html/template"
	"strconv"
	"strings"
	"unicode/utf8"
)

// QueryArticleById 依据 id 查询文章
func QueryArticleById(id uint64) (entity.Article, error) {
	var article entity.Article
	err := dal.DB.Where("id = ?", id).Take(&article).Error
	return article, err
}

// QueryArticlesByUid 依据 uid 查询所有文章
func QueryArticlesByUid(uid uint64) ([]entity.Article, error) {
	var articles []entity.Article
	err := dal.DB.Where("id = ?", uid).Find(&articles).Error
	return articles, err
}

// QueryArticleCountByUid 依据 uid 查询文章总数
func QueryArticleCountByUid(uid uint64) (int, error) {
	var count int
	err := dal.DB.Model(entity.Article{}).Where("userid = ?", uid).Count(&count).Error
	return count, err
}

// LimitQueryOrderedArticlesByUid 依据 uid 分页查询文章
// orderCol: 按照那一列排序
// order: 顺序 false 为 asc，true 为 desc
func LimitQueryOrderedArticlesByUid(uid uint64, index int, size int, orderCol string, orderBy bool) ([]entity.Article, error) {
	// 正序 or 倒叙
	var status string
	if !orderBy {
		// 默认正序
		status = orderCol + " asc"
	} else {
		// 倒叙
		status = orderCol + " desc"
	}
	// 文章切片
	var articles []entity.Article
	err := dal.DB.Where("userid = ?", uid).Limit(size).Offset(index).Order(status).Find(&articles).Error
	return articles, err
}

// AddArticle 新增一片文章
func AddArticle(article entity.Article) error {
	return dal.DB.Create(&article).Error
}

// UpdateArticleById 依据 id 更新文章后
func UpdateArticleById(article entity.Article, id uint64) error {
	return dal.DB.Where("id = ?", id).Update(article).Error
}

// DeleteArticleById 依据 id 删除文章
func DeleteArticleById(id uint64) error {
	return dal.DB.Where("id = ?", id).Delete(&entity.Article{}).Error
}

/* ================ 基础操作分割线 ================ */

// MakeUserArticleListTemplate 依据用户 id 和页码创建文章列表模版
func MakeUserArticleListTemplate(uid uint64, pageNum int) (template.HTML, error) {
	if pageNum <= 0 {
		pageNum = 1
	}

	// 对应 pageNum 页首行文章下标
	index := (pageNum - 1) * config.SIZE_PERPAGE
	// 每页 size
	size := config.SIZE_PERPAGE
	// order column
	orderCol := "updated_at"

	// 分页查询数据库获取文章列表 desc
	articles, err := LimitQueryOrderedArticlesByUid(uid, index, size, orderCol, true)
	if err != gorm.ErrRecordNotFound && err != nil {
		return "", errors.New("分页查询出错导致模版为空")
	}

	var articlePage string
	// 遍历文章切片
	for _, article := range articles {
		// 修改 article content 内容加上 ...
		if utf8.RuneCountInString(article.Content) > config.CHAR_PERARTICLE {
			article.Content = string([]rune(article.Content)[0:100]) + "..."
		}
		// format Tags
		article.Tags = ChangeTags(article.Tags)
		// 解析模版文件
		tem, err := template.ParseFiles("view/article/article_list.html")
		if err != nil {
			return "", err
		}
		// 解析后的模版插入变量到 buff
		var buff bytes.Buffer
		err = tem.Execute(&buff, map[string]interface{}{
			"Id":        article.ID,
			"Title":     article.Title,
			"Tags":      template.HTML(article.Tags),
			"Short":     article.Short,
			"Content":   article.Content,
			"Uid":       uid,
			"CreatedAt": article.CreatedAt,
			"UpdatedAt": article.UpdatedAt,
		})
		if err != nil {
			return "", err
		}
		// buff => articlePage
		articlePage += buff.String()
	}
	return template.HTML(articlePage), nil
}

// ChangeTags 将 Tags 修改成指定标签格式可被浏览器识别
func ChangeTags(tags string) string {
	// 判定文章类型
	switch tags {
	// 程序博客
	case config.PROGRAM:
		tags = "<span class=\"label label-warning\">程序博客</span>"
	// 金融理财
	case config.FINANCE:
		tags = "<span class=\"label label-primary\">金融理财</span>"
	// 自然科学
	case config.SCIENCE:
		tags = "<span class=\"label label-success\">自然科学</span>"
	// 音乐艺术
	case config.ART:
		tags = "<span class=\"label label-info\">音乐艺术</span>"
	// 体育竞技
	case config.SPORT:
		tags = "<span class=\"label label-danger\">体育竞技</span>"
	// 其他
	default:
		tags = "<span class=\"label label-default\">其他</span>"
	}
	return tags
}

// MakePageBarArgs 创建页码参数
func MakePageBarArgs(uid uint64, pageNum int) (map[string]interface{}, error) {
	// 拿到文章总数
	allNum, err := QueryArticleCountByUid(uid)
	if err != gorm.ErrRecordNotFound && err != nil {
		return map[string]interface{}{}, err
	}
	if pageNum <= 0 {
		pageNum = 1
	}

	pageBarMap := make(map[string]interface{})
	// 当前页码数
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
	return pageBarMap, nil
}

// ChangeNToBr 转换 content 为 <div> 可识别的换行符号
func ChangeNToBr(article entity.Article) entity.Article {
	// 将文本内容替换成可以显示换行的文本
	article.Content = strings.Replace(article.Content, "\r\n", "<br>", -1)
	article.Content = strings.Replace(article.Content, "\n", "<br>", -1)
	return article
}

// ChangeNToNextLine 转换 content 为 text 中可识别换行符
func ChangeNToNextLine(article entity.Article) entity.Article {
	// 格式化 content
	article.Content = strings.Replace(article.Content, "\r\n", "&#13;&#10;", -1)
	article.Content = strings.Replace(article.Content, "\n", "&#13;&#10;", -1)
	return article
}
