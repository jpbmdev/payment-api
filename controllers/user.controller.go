package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/services"
)

// -----------------------------------------------
// -- User controller
// -----------------------------------------------
type UserController interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
}

type userController struct {
	service services.UserService
}

//Function to crete new user controller
func NewUserController() UserController {
	return &userController{
		service: services.NewUserService(),
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
		ctx.JSON(http.StatusBadRequest, models.FailedOperation{
			InternalCode: "BadRequest",
			Message:      err.Error(),
		})
		return
	}

	//Create user model from dto
	user := models.User{Name: createUserDto.Name}

	//Create user in database
	err = c.service.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.FailedOperation{
			InternalCode: "InternalServerError",
			Message:      err.Error(),
		})
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
		ctx.JSON(http.StatusInternalServerError, models.FailedOperation{
			InternalCode: "InternalServerError",
			Message:      err.Error(),
		})
		return
	}

	//Return list of users
	ctx.JSON(http.StatusOK, users)
}
