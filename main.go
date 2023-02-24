package main

import (
	"flag"
	internal "gateway/lib/init"
	"gateway/models"
	"gateway/proxy/http_proxy/http_proxy_router"
	"gateway/service/router"
	"gateway/web"
	"os"
	"os/signal"
	"syscall"
)

func parseArgs() (*string, *string) {
	var (
		endpoint = flag.String("endpoint", "", "required: input endpoint server/static")
		config   = flag.String("config", "./conf/config.yml", "option: config file path")
	)
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}

	return endpoint, config
}

func main() {
	endpoint, config := parseArgs()

	internal.InitModule(*config)
	if *endpoint == "server" {
		models.ServiceManagerHandler.LoadOnce()
		go func() {
			router.HttpServerRun()
		}()

		go func() {
			http_proxy_router.HttpServerRun()
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		http_proxy_router.HttpServerStop()
		router.HttpServerStop()
	} else if *endpoint == "static" {
		web.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		web.HttpServerStop()
	}
}
