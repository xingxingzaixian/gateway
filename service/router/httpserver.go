package router

import (
	"gateway/lib/config"
	"gateway/service/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(config.BaseConfig.DebugMode)
	r := InitRouter(middleware.Cors(), middleware.TranslationMiddleware())
	HttpSrvHandler = &http.Server{
		Addr:           config.BaseConfig.Http.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(config.BaseConfig.Http.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.BaseConfig.Http.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(config.BaseConfig.Http.MaxHeaderBytes),
	}
	go func() {
		zap.S().Infof(" [INFO] HttpServerRun:%s\n", config.BaseConfig.Http.Addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			zap.S().Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", config.BaseConfig.Http.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		zap.S().Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	zap.S().Infof(" [INFO] HttpServerStop stopped\n")
}
