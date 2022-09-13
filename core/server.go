package core

import (
	"os"

	"github.com/gin-gonic/gin"
)

type ServerInstance interface {
	InitServer()
}

type serverInstace struct {
	gin *gin.Engine
}

func NewServerInstance() ServerInstance {
	return &serverInstace{
		gin: gin.Default(),
	}
}

func (s *serverInstace) InitServer() {

	s.gin.Run(os.Getenv("PORT"))
}
