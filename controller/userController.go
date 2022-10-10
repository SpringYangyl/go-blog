package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/dao"
	"go-blog/response"
	"go-blog/service"
)

func Register(c *gin.Context) {
	json := UserRequest{}
	c.BindJSON(&json)

	username := json.Username
	password := json.Password
	user := &dao.User{
		Username: username,
		Password: password,
		Role:     "normal",
	}
	if service.Register(user) {
		response.ResponseSuccess(c, "注册成功", true)
	} else {
		response.ResponseError(c, "注册失败")
	}
}
func Login(c *gin.Context) {
	json := UserRequest{}
	c.BindJSON(&json)

	username := json.Username
	password := json.Password
	res := service.Login(username, password)
	if res {
		response.ResponseSuccess(c, "登录成功", true)

	} else {
		response.ResponseError(c, "登录失败")
	}
}

type UserRequest struct {
	Username string
	Password string
}
