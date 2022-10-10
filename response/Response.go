package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(msg string, data any) any {
	return gin.H{
		"msg":  msg,
		"data": data,
	}
}
func ResponseSuccess(c *gin.Context, msg string, data any) {
	c.JSON(200, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
}
func ResponseError(c *gin.Context, msg string) {

	c.JSON(400, gin.H{
		"code": 400,
		"msg":  msg,
		"data": nil,
	})

}
