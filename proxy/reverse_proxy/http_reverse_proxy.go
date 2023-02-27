package reverse_proxy

import (
	"gateway/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func NewHttpReverseProxy(c *gin.Context, serviceDetail *models.ServiceDetail) *httputil.ReverseProxy {
	//请求协调者
	director := func(req *http.Request) {
		nextAddr := serviceDetail.HTTPRule.UrlRewrite
		//todo 优化点3
		if nextAddr == "" {
			panic("get next addr fail")
		}
		target, err := url.Parse(nextAddr)
		if err != nil {
			panic(err)
		}

		targetQuery := target.RawQuery
		zap.S().Infof("before rewrite: %s", req.URL.String())
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = strings.Replace(req.URL.Path, serviceDetail.HTTPRule.Rule, target.Path, 1)
		req.Host = target.Host
		zap.S().Infof("after rewrite: %s", req.URL.String())
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			req.Header.Set("User-Agent", "user-agent")
		}
	}

	//更改内容
	modifyFunc := func(resp *http.Response) error {
		// websocket服务不需要处理
		if strings.Contains(resp.Header.Get("Connection"), "Upgrade") {
			return nil
		}
		return nil
	}

	//错误回调 ：关闭real_server时测试，错误回调
	//范围：transport.RoundTrip发生的错误、以及ModifyResponse发生的错误
	errFunc := func(w http.ResponseWriter, r *http.Request, err error) {
		//middleware.ResponseError(c, 999, err)
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc, ErrorHandler: errFunc}
}
