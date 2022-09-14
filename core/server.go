package core

import (
	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/config"
	"github.com/jpbmdev/payment-api/routes"
)

// -----------------------------------------------
// -- Struct to create and configure the server
// -----------------------------------------------
type ServerInstance interface {
	InitServer()
}

type serverInstace struct {
	gin                *gin.Engine
	userRoutes         routes.UserRoutes
	targetSchemaRoutes routes.TargetShemaRoutes
}

//Function to crete new server instance
func NewServerInstance() ServerInstance {
	return &serverInstace{
		gin:                gin.Default(),
		userRoutes:         routes.NewUserRoutes(),
		targetSchemaRoutes: routes.NewTargetShemaRoutes(),
	}
}

func (s *serverInstace) InitServer() {
	//Load Swagger documentation page
	InitializeSwagger(s.gin)

	//Initialize user routes
	s.userRoutes.InitializeRoutes(s.gin)

	//Initialize targetSchema routes
	s.targetSchemaRoutes.InitializeRoutes(s.gin)

	//Initialize server
	s.gin.Run(config.ConfigSchema.Port)
}
