package http_proxy_router

import (
	"gateway/proxy/http_proxy/http_proxy_middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Use(
		http_proxy_middleware.HTTPAccessModeMiddleware(),
		http_proxy_middleware.HTTPReverseProxyMiddleware(),
	)

	return router
}
