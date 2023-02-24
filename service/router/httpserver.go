package router

import (
	"gateway/lib/config"
	"gateway/service/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
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
		log.Printf(" [INFO] HttpServerRun:%s\n", config.BaseConfig.Http.Addr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", config.BaseConfig.Http.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
