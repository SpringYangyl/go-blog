package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/controller"
	"go-blog/middleware/auth"
)

func initRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/jwt/", auth.Auth(), controller.Jwt)
	v1.POST("/register/", controller.Register)
	v1.POST("/login/", controller.Login)
	v1.GET("/getArticleById/", controller.GetArticleContent)
	v1.GET("/getArticleList", controller.GetAllArticle)
	v1.POST("/createArticle/", controller.CreateArticle)
}
