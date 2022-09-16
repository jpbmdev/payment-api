package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/controllers"
	"github.com/jpbmdev/payment-api/middlewares"
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
	server.GET("/loan", middlewares.PaginationMiddleware(), r.controller.GetLoans)
	server.POST("/loan", r.controller.CreateLoan)
	server.GET("/loan/:id", r.controller.GetLoanById)
	server.GET("/loan/:id/debt", r.controller.GetLoanDebt)
	server.GET("/loan/:id/payment", r.controller.GetPaymentsByLoanId)
	server.PUT("/loan/:id/payment", r.controller.AddPaymentToLoan)
}
