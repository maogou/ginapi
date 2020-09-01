package model

import "github.com/maogou/ginapi/pkg/app"

type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (t Tag)TableName() string{
	return "blog_tag"
}

//针对swagger生成文档包含别的结构体无法显示的bug
type TagSwagger struct {
	List []*Tag
	Pager *app.Page
}
