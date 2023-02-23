package web

import (
	"gateway/lib/config"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"log"
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
		log.Printf(" [INFO] HttpStaticRun:%s\n", config.BaseConfig.Web.Addr)
		if err := HttpStaticHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpStaticRun:%s err:%v\n", config.BaseConfig.Web.Addr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpStaticHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
