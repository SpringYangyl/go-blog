package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/controller"
	"go-blog/middleware"
)

func initRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	v1.GET("/jwt/", middleware.Auth(), controller.Jwt)
	v1.POST("/register/", controller.Register)
	v1.POST("/login/", controller.Login)
}
