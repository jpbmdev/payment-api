package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/controllers"
)

// -----------------------------------------------
// -- Struct to add targetShema routes to gin router
// -----------------------------------------------
type TargetShemaRoutes interface {
	InitializeRoutes(server *gin.Engine)
}

type targetShemaRoutes struct {
	controller controllers.TargetSchemaController
}

//Function to crete new targetSchema routes
func NewTargetShemaRoutes() TargetShemaRoutes {
	return &targetShemaRoutes{
		controller: controllers.NewTargetSchemaController(),
	}
}

func (r *targetShemaRoutes) InitializeRoutes(server *gin.Engine) {
	server.POST("/target-schema/test-tree", r.controller.TestTargetSchemaDecisionTree)
	server.GET("/target-schema", r.controller.GetTargetSchemas)
}
