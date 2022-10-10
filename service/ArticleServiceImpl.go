package service

import "go-blog/dao"

func CreateArticle(article dao.Article) bool {
	return dao.CreateArticle(article)
}
func GetAllArticle() ([]dao.Article, error) {
	return dao.GetAllArticle()
}
func GetArticleContent(id int64) (dao.Article, error) {
	return dao.GetArticleContent(id)
}
