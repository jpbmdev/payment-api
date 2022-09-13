package core

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jpbmdev/payment-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Payment Simple API
// @version         1.0
// @description     This is a sample server to manage payments
// @host      localhost:8080
// @BasePath  /
func InitializeSwagger(server *gin.Engine) {
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
