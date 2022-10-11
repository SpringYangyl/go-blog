package main

import (
	"github.com/gin-gonic/gin"
	"go-blog/dao"
	"go-blog/middleware"
)

func main() {
	middleware.Init()
	dao.Init()
	router := gin.Default()
	initRouter(router)
	router.Run(":8080")

}
