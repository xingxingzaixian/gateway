package http_proxy_router

import (
	"gateway/lib/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

var (
	HttpSrvHandler  *http.Server
	HttpsSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(config.BaseConfig.DebugMode)
	router := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           config.GetStringConf("proxy.http.addr"),
		Handler:        router,
		ReadTimeout:    time.Duration(config.GetIntConf("proxy.http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(config.GetIntConf("proxy.http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(config.GetIntConf("proxy.http.max_header_bytes")),
	}

	zap.S().Infof(" [INFO] http_proxy_run %s\n", config.GetStringConf("proxy.http.addr"))
	if err := HttpSrvHandler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.S().Fatalf(" [ERROR] http_proxy_run %s err:%v\n", config.GetStringConf("proxy.http.addr"), err)
	}
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		zap.S().Errorf(" [ERROR] http_proxy_stop err:%v\n", err)
	}
	zap.S().Infof(" [INFO] http_proxy_stop %v stopped\n", config.GetStringConf("proxy.http.addr"))
}

func HttpsServerRun() {
	gin.SetMode(config.BaseConfig.DebugMode)
	r := InitRouter()
	HttpsSrvHandler = &http.Server{
		Addr:           config.GetStringConf("proxy.https.addr"),
		Handler:        r,
		ReadTimeout:    time.Duration(config.GetIntConf("proxy.https.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(config.GetIntConf("proxy.https.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(config.GetIntConf("proxy.https.max_header_bytes")),
	}
	zap.S().Infof(" [INFO] https_proxy_run %s\n", config.GetStringConf("proxy.https.addr"))
	//todo 以下命令只在编译机有效，如果是交叉编译情况下需要单独设置路径
	if err := HttpsSrvHandler.ListenAndServeTLS(config.GetStringConf("proxy.https.cert_crt_file"), config.GetStringConf("proxy.https.cert_key_file")); err != nil && err != http.ErrServerClosed {
		zap.S().Errorf(" [ERROR] https_proxy_run %s err:%v\n", config.GetStringConf("proxy.https.addr"), err)
	}
}

func HttpsServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpsSrvHandler.Shutdown(ctx); err != nil {
		zap.S().Fatalf(" [ERROR] https_proxy_stop err:%v\n", err)
	}
	zap.S().Infof(" [INFO] https_proxy_stop %v stopped\n", config.GetStringConf("proxy.https.addr"))
}
