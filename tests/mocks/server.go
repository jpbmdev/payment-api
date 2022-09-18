package mocks_test

import (
	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/config"
)

func GetRouter() *gin.Engine {
	//Load enviroment variables
	config.LoadConfig()

	r := gin.Default()
	return r
}
