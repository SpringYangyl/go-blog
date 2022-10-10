package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Jwt(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}
