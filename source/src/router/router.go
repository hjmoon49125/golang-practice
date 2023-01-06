package router

import (
	"server_temp/source/src/controller"
	"server_temp/source/src/middleware"

	"github.com/gin-gonic/gin"

	_ "server_temp/source/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoute() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/auth")
	controller.AuthController(authGroup)

	apiGroup := router.Group("/api", middleware.JWTAuthMiddleware())
	controller.HelloController(apiGroup)

	if mode := gin.Mode(); mode == gin.DebugMode {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
