package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/dao"
	"go-blog/response"
	"go-blog/service"
	"strconv"
)

func CreateArticle(c *gin.Context) {
	json := ArticleRequest{}
	c.BindJSON(&json)
	name := json.ArticleName
	userId := json.UserId
	desc := json.Description
	content := json.Content

	article := dao.Article{
		ArticleName: name,
		UserId:      userId,
		Description: desc,
		Content:     content,
		WatchCount:  int64(0),
		LikeCount:   int64(0),
	}
	if service.CreateArticle(article) {
		response.ResponseSuccess(c, "创建成功", true)

	} else {
		response.ResponseError(c, "创建失败")
	}
}

func GetAllArticle(c *gin.Context) {
	articles, err := service.GetAllArticle()
	if err != nil {
		response.ResponseError(c, "系统错误")
	} else {
		response.ResponseSuccess(c, "success", articles)
	}
}

func GetArticleContent(c *gin.Context) {
	id := c.Query("id")
	articleId, errs := strconv.ParseInt(id, 10, 64)
	if errs != nil {
		response.ResponseError(c, "失败")
	}
	article, err := service.GetArticleContent(articleId)
	if err != nil {
		response.ResponseError(c, "系统错误")
	} else {
		response.ResponseSuccess(c, "成功", article)
	}
}

type ArticleRequest struct {
	UserId      int64  `json:"user_id"`
	ArticleName string `json:"article_name"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
