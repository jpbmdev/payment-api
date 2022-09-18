package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/controllers"
)

// -----------------------------------------------
// -- Struct to add user routes to gin router
// -----------------------------------------------
type UserRoutes interface {
	InitializeRoutes(server *gin.Engine)
}

type userRoutes struct {
	controller controllers.UserController
}

//Function to crete new user routes
func NewUserRoutes() UserRoutes {
	return &userRoutes{
		controller: controllers.NewUserController(),
	}
}

func (r *userRoutes) InitializeRoutes(server *gin.Engine) {
	server.POST("/user", r.controller.CreateUser)
	server.GET("/user", r.controller.GetUsers)
	server.GET("/user/:id/loan", r.controller.GetUserLoans)
}
