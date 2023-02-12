package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pay/middleware"
	payRouter "pay/router"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	apiRouter := router.Group("")
	payRouter.InitPayRouter(apiRouter)

	return router
}
