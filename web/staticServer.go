package web

import (
	"gateway/lib/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

var (
	HttpStaticHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode(config.BaseConfig.DebugMode)
	r := InitRouter()
	HttpStaticHandler = &http.Server{
		Addr:           config.BaseConfig.Web.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(config.BaseConfig.Web.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.BaseConfig.Web.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << uint(config.BaseConfig.Web.MaxHeaderBytes),
	}
	go func() {
		zap.S().Infof(" [INFO] HttpStaticRun:%s\n", config.BaseConfig.Web.Addr)
		if err := HttpStaticHandler.ListenAndServe(); err != nil {
			zap.S().Fatalf(" [ERROR] HttpStaticRun:%s err:%v\n", config.BaseConfig.Web.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpStaticHandler.Shutdown(ctx); err != nil {
		zap.S().Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	zap.S().Infof(" [INFO] HttpServerStop stopped\n")
}
