package controller

import (
	"github.com/gin-gonic/gin"
	"go-blog/dao"
	"go-blog/service"
	"net/http"
)

func Jwt(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := &dao.User{
		Username: username,
		Password: password,
		Role:     "normal",
	}
	if service.Register(user) {
		c.JSON(http.StatusOK, "注册成功")
	}
	c.JSON(400, "注册失败")
}
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	res := service.Login(username, password)
	if res {
		c.JSON(http.StatusOK, "登录成功")
	}
	c.JSON(400, "登录失败")
}
