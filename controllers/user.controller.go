package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/errorsResponse"
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// -----------------------------------------------
// -- User controller
// -----------------------------------------------
type UserController interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	GetUserLoans(ctx *gin.Context)
}

type userController struct {
	service     services.UserService
	loanService services.LoanService
}

//Function to crete new user controller
func NewUserController() UserController {
	return &userController{
		service:     services.NewUserService(),
		loanService: services.NewLoanService(),
	}
}

// CreateUser godoc
// @Summary Create User
// @Schemes
// @Description Create a user with his name
// @Tags user
// @Accept json
// @Produce json
// @Param CreateUserDto body models.CreateUserDto true "payload"
// @Success 201 {object}  models.SucessfullOperation
// @Failure 400 {object}  models.FailedOperation
// @Failure 500 {object}  models.FailedOperation
// @Router /user [post]
func (c *userController) CreateUser(ctx *gin.Context) {
	var createUserDto models.CreateUserDto

	//Get request body
	err := ctx.ShouldBindJSON(&createUserDto)

	if err != nil {
		errorsResponse.Error400(ctx, err.Error())
		return
	}

	//Create user model from dto
	user := models.User{Name: createUserDto.Name}

	//Create user in database
	err = c.service.CreateUser(user)

	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Create Success Response
	ctx.JSON(http.StatusCreated, models.SucessfullOperation{Message: "Exito"})
}

// GetUsers godoc
// @Summary Get users
// @Schemes
// @Description Get lists of users
// @Tags user
// @Produce json
// @Success 200 {array}   models.User
// @Failure 500 {object}  models.FailedOperation
// @Router /user [get]
func (c *userController) GetUsers(ctx *gin.Context) {

	//Get list of users from db
	users, err := c.service.GetUsers()

	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Return list of users
	ctx.JSON(http.StatusOK, users)
}

// GetUserLoans godoc
// @Summary Get user loans
// @Schemes
// @Description Get lists user loans
// @Tags user
// @Produce json
// @Param id  path string true "ID"
// @Success 200 {array}   models.Loans
// @Failure 400 {object}  models.FailedOperation
// @Failure 500 {object}  models.FailedOperation
// @Router /user/{id}/loan [get]
func (c *userController) GetUserLoans(ctx *gin.Context) {

	//Check if the id passed is a mongoID
	userId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		errorsResponse.Error400(ctx, "Invalid User ID")
		return
	}

	//Get list of loans from db
	loans, err := c.loanService.FindLoansByUserId(userId)

	if err != nil {
		errorsResponse.Error500(ctx, err.Error())
		return
	}

	//Return list of loans
	ctx.JSON(http.StatusOK, loans)
}
