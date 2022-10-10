package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/dao"
	"go-blog/service"
	"net/http"
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
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册成功",
		})
	} else {
		c.JSON(400, gin.H{
			"msg": "注册失败",
		})

	}
}
func Login(c *gin.Context) {
	json := UserRequest{}
	c.BindJSON(&json)

	username := json.Username
	password := json.Password
	res := service.Login(username, password)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录成功",
		})
	} else {
		c.JSON(400, gin.H{
			"msg": "登录失败",
		})
	}
}

type UserRequest struct {
	Username string
	Password string
}
