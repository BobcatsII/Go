package model

import "gin-blog/pkg/app"

type Tag struct {
	*Model        // 将Model结构体引入
	Name   string `json:"name"`
	State  uint8  `json:"state"`
}

type TagSwagger struct {
	List []*Tag
	Pager *app.Pager
}

func (a Tag) TableName() string {
	return "blog_tag"
}

