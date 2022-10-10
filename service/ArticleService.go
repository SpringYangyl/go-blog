package service

import "go-blog/dao"

type ArticleService interface {
	CreateArticle(article dao.Article) bool
	GetAllArticle() ([]dao.Article, error)
	GetArticleContent(id int64) (dao.Article, error)
}
