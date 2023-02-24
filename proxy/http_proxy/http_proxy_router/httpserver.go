package http_proxy_router

import (
	"gateway/lib/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

var HttpSrvHandler *http.Server

func HttpServerRun() {
	gin.SetMode(config.BaseConfig.DebugMode)
	router := InitRouter()
	HttpSrvHandler = &http.Server{
		Addr:           config.BaseConfig.Proxy.Addr,
		Handler:        router,
		ReadTimeout:    time.Duration(config.BaseConfig.Proxy.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.BaseConfig.Proxy.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(config.BaseConfig.Proxy.MaxHeaderBytes),
	}

	zap.S().Infof(" [INFO] http_proxy_run %s\n", config.BaseConfig.Proxy.Addr)
	if err := HttpSrvHandler.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zap.S().Fatalf(" [ERROR] http_proxy_run %s err:%v\n", config.BaseConfig.Proxy.Addr, err)
	}
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		zap.S().Infof(" [ERROR] http_proxy_stop err:%v\n", err)
	}
	zap.S().Infof(" [INFO] http_proxy_stop %v stopped\n", config.BaseConfig.Proxy.Addr)
}
