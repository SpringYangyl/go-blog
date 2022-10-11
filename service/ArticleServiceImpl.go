package service

import (
	"fmt"
	"go-blog/dao"
	"go-blog/middleware"
)

var redisPrefix = "article:"

func CreateArticle(article dao.Article) bool {
	return dao.CreateArticle(article)
}
func GetAllArticle() ([]dao.Article, error) {
	articles, err := dao.GetAllArticle()
	if err != nil {
		return articles, err
	}
	for _, val := range articles {
		count := GetWatchCountFromRedis(val.Id)
		if count == 0 {
			continue
		}
		val.WatchCount = count
	}
	return articles, nil
}
func GetArticleContent(id int64) (dao.Article, error) {
	article, err := dao.GetArticleContent(id)
	if err != nil {
		return article, err
	}
	count := GetWatchCountFromRedis(article.Id)
	if count != 0 {
		article.WatchCount = count
	}
	return article, nil

}
func GetWatchCountFromRedis(id int64) int64 {
	res, err := middleware.GetKey(redisPrefix + string(id))
	if err != nil {
		return int64(0)
	}
	return res.(int64)
}
func LikeArticle(id int64) bool {
	article, err := dao.GetArticleContent(id)
	if err != nil {
		return false
	}

	fmt.Println(article)
	return true
}
