package model

import "github.com/maogou/ginapi/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article)TableName() string  {
	return "blog_article"
}

//针对swagger生成文档包含别的结构体无法显示的bug
type ArticleSwagger struct {
	List []*Article
	pager *app.Page
}
