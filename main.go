package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/dao"
	"go-blog/middleware"
	"go-blog/service"
)

func main() {
	middleware.Init()
	dao.Init()
	user,_ := dao.GetUserByName("zhangsan")
	res :=service.NewToken(user)
	middleware.SetToken(res)
	fmt.Println(middleware.GetToken())
	router := gin.Default()
	initRouter(router)


	router.Run(":8080")
	//user := &dao.User{
	//	Username: "zhangsan",
	//	Password: "123456",
	//	Role:     "admin",
	//}
	//service.Register(user)
}

