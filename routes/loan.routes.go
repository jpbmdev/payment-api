package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/controllers"
)

// -----------------------------------------------
// -- Struct to add loan routes to gin router
// -----------------------------------------------
type LoanRoutes interface {
	InitializeRoutes(server *gin.Engine)
}

type loanRoutes struct {
	controller controllers.LoanController
}

//Function to crete new user routes
func NewLoanRoutes() LoanRoutes {
	return &loanRoutes{
		controller: controllers.NewLoanController(),
	}
}

func (r *loanRoutes) InitializeRoutes(server *gin.Engine) {
	server.POST("/loan", r.controller.CreateLoan)
}
