package router

import (
	"github.com/gin-gonic/gin"
	"pay/api/ali"
	"pay/middleware"
)

func InitPayRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("pay/").Use(middleware.Cors()).Use(middleware.LoggerToFile())
	{
		MessageRouter.POST("alipay/notify", ali.Notify)
		MessageRouter.Use(middleware.Auth()).POST("alipay/web", ali.Web)
	}
}
