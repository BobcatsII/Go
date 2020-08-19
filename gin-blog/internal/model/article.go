package model

import "gin-blog/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	CoverImageUrl string `json:"cover_image_url"`
	Content       string `json:"content"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

