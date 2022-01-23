package model

type ArticleGetReq struct {
	PageNum int `form:"page"`
}

type ArticleChangeReq struct {
	Title   string `form:"title"`
	Tags    string `form:"tags"`
	Short   string `form:"short"`
	Content string `form:"content"`
}
