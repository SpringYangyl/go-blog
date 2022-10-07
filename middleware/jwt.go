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
			context.JSON(400,http.Response{
				Status:           "错误",
				StatusCode:       0,
				Proto:            "",
				ProtoMajor:       0,
				ProtoMinor:       0,
				Header:           nil,
				Body:             nil,
				ContentLength:    0,
				TransferEncoding: nil,
				Close:            false,
				Uncompressed:     false,
				Trailer:          nil,
				Request:          nil,
				TLS:              nil,
			})
		}
		if parseToken.Issuer != "admin"{
			context.JSON(401,http.Response{
				Status:           "无权限",
				StatusCode:       0,
				Proto:            "",
				ProtoMajor:       0,
				ProtoMinor:       0,
				Header:           nil,
				Body:             nil,
				ContentLength:    0,
				TransferEncoding: nil,
				Close:            false,
				Uncompressed:     false,
				Trailer:          nil,
				Request:          nil,
				TLS:              nil,
			})
		}
		context.JSON(http.StatusOK,http.Response{
			Status: string("欢迎管理员"),
		})
	}
	
}
