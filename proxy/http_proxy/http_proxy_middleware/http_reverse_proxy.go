package http_proxy_middleware

import (
	"gateway/models"
	"gateway/proxy/http_proxy/reverse_proxy"
	"gateway/public"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 匹配接入方式 基于请求信息
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			public.ResponseError(c, public.MiddleReverseProxy, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*models.ServiceDetail)

		proxy := reverse_proxy.NewReverseProxy(c, serviceDetail)
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return
	}
}
