package http_proxy_router

import (
	"gateway/cert_file"
	"gateway/lib/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"log"
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
		Addr:           config.BaseConfig.ProxyHttp.Addr,
		Handler:        router,
		ReadTimeout:    time.Duration(config.BaseConfig.ProxyHttp.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.BaseConfig.ProxyHttp.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(config.BaseConfig.ProxyHttp.MaxHeaderBytes),
	}

	zap.S().Infof(" [INFO] http_proxy_run %s\n", config.BaseConfig.ProxyHttp.Addr)
	if err := HttpSrvHandler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.S().Fatalf(" [ERROR] http_proxy_run %s err:%v\n", config.BaseConfig.ProxyHttp.Addr, err)
	}
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		zap.S().Infof(" [ERROR] http_proxy_stop err:%v\n", err)
	}
	zap.S().Infof(" [INFO] http_proxy_stop %v stopped\n", config.BaseConfig.ProxyHttp.Addr)
}

func HttpsServerRun() {
	gin.SetMode(config.BaseConfig.DebugMode)
	r := InitRouter()
	HttpsSrvHandler = &http.Server{
		Addr:           config.BaseConfig.ProxyHttps.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(config.BaseConfig.ProxyHttps.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.BaseConfig.ProxyHttps.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(config.BaseConfig.ProxyHttps.MaxHeaderBytes),
	}
	log.Printf(" [INFO] https_proxy_run %s\n", config.BaseConfig.ProxyHttps.Addr)
	//todo 以下命令只在编译机有效，如果是交叉编译情况下需要单独设置路径
	if err := HttpsSrvHandler.ListenAndServeTLS(cert_file.Path("server.crt"), cert_file.Path("server.key")); err != nil && err != http.ErrServerClosed {
		//if err := HttpsSrvHandler.ListenAndServeTLS("./cert_file/server.crt", "./cert_file/server.key"); err != nil && err != http.ErrServerClosed {
		log.Fatalf(" [ERROR] https_proxy_run %s err:%v\n", config.BaseConfig.ProxyHttps.Addr, err)
	}
}

func HttpsServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpsSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] https_proxy_stop err:%v\n", err)
	}
	log.Printf(" [INFO] https_proxy_stop %v stopped\n", config.BaseConfig.ProxyHttps.Addr)
}
