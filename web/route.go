package web

import (
	"gateway/lib/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.Request.URL.Path = config.BaseConfig.Web.Url
		router.HandleContext(ctx)
	})

	// 将静态服务器配置在此
	router.Static(config.BaseConfig.Web.Url, config.BaseConfig.Web.Path)

	return router
}
