package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/dao"
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
		c.JSON(200, gin.H{
			"msg": "创建成功",
		})
	} else {
		c.JSON(400, gin.H{
			"msg": "创建失败",
		})

	}
}

func GetAllArticle(c *gin.Context) {
	articles, err := service.GetAllArticle()
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "系统错误",
		})
	} else {
		c.JSON(200, gin.H{
			"data": articles,
		})
	}
}

func GetArticleContent(c *gin.Context) {
	id := c.Query("id")
	articleId, errs := strconv.ParseInt(id, 10, 64)
	if errs != nil {
		c.JSON(400, "系统错误")
	}
	articlc, err := service.GetArticleContent(articleId)
	if err != nil {
		c.JSON(400, "查询失败")
	}
	c.JSON(200, articlc)
}

type ArticleRequest struct {
	UserId      int64  `json:"user_id"`
	ArticleName string `json:"article_name"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
