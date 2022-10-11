package auth

import (
	"github.com/gin-gonic/gin"
	"go-blog/response"
	"go-blog/service"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if token == "" {
			response.ResponseError(context, "NO token available")
		}
		parseToken, err := service.ParseToken(token)
		if err != nil {
			response.ResponseError(context, "token Not Correct")
		}

		if parseToken.Issuer != "admin" {
			response.ResponseError(context, "user have no power")
		}
		response.ResponseSuccess(context, "welcome admin !", nil)

	}

}
