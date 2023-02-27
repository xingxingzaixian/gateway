package router

import (
	"gateway/docs"
	"gateway/lib/config"
	"gateway/service/api"
	"gateway/service/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	docs.SwaggerInfo.Title = config.BaseConfig.Swagger.Title
	docs.SwaggerInfo.Description = config.BaseConfig.Swagger.Desc
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.BaseConfig.Swagger.Host
	docs.SwaggerInfo.BasePath = config.BaseConfig.Swagger.BasePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	serviceRouter := router.Group("/service")
	{
		api.ServiceRegister(serviceRouter)
	}

	adminLoginRouter := router.Group("/admin_login")
	{
		api.AdminLoginRegister(adminLoginRouter)
	}

	adminRouter := router.Group("/admin")
	adminRouter.Use(middleware.JwtAuth())
	{
		api.AdminRegister(adminRouter)
	}
	return router
}
