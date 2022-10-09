package middleware

import (
	"github.com/gin-gonic/gin"
	"go-blog/service"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		//auth := context.Query("token")
		//if len(auth) == 0{
		//	context.JSON(http.StatusUnauthorized,http.Response{
		//		Status: string("没有token"),
		//	})
		//}
		//token,err := service.ParseToken(auth)
		//if err != nil{
		//	context.JSON(http.StatusUnauthorized,http.Response{Status: string(err.Error())})
		//}
		//if token.Issuer != "admin"{
		//	context.JSON(http.StatusUnauthorized,http.Response{
		//		Status: string("没有权限"),
		//	})
		//}
		token := GetToken()
		parseToken, err := service.ParseToken(token)
		if err != nil {
			context.JSON(400, "错误")
		}
		if parseToken.Issuer != "admin" {
			context.JSON(401, "无权限")
		}
		context.JSON(http.StatusOK, "欢迎管理员")

	}

}
