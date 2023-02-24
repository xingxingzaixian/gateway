package reverse_proxy

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReverseProxy interface {
	ServeHTTP(rw http.ResponseWriter, req *http.Request)
}

func NewReverseProxy(c *gin.Context, serviceDetail *models.ServiceDetail) ReverseProxy {
	if serviceDetail.HTTPRule.NeedWebsocket == 1 {
		return NewWebsocketReverseProxy(c, serviceDetail)
	} else {
		return NewHttpReverseProxy(c, serviceDetail)
	}
}
