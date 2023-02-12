package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"pay/api"
	userProto "pay/api/qvbilam/user/v1"
	"pay/global"
)

// Auth 验证jwt
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := global.UserServerClient.Auth(context.Background(), &userProto.AuthRequest{
			Token: ctx.Request.Header.Get("Authorization"),
		})

		if err != nil {
			fmt.Println(ctx.Request.Header.Get("Authorization"))
			fmt.Println("让你登录呢")
			api.HandleValidateError(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Set("userId", user.Id)
		// 继续执行
		ctx.Next()
	}
}