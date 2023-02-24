package http_proxy_middleware

import (
	"gateway/models"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 匹配接入方式 基于请求信息
func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := models.ServiceManagerHandler.HTTPAccessMode(c)
		if err != nil {
			public.ResponseError(c, public.MiddleAccessMode, err)
			zap.S().Error(err)
			c.Abort()
			return
		}

		c.Set("service", service)
		c.Next()
	}
}
