package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/models"
	"github.com/jpbmdev/payment-api/services"
)

// -----------------------------------------------
// -- TargetSchema controller
// -----------------------------------------------
type TargetSchemaController interface {
	GetTargetSchemas(ctx *gin.Context)
	TestTargetSchemaDecisionTree(ctx *gin.Context)
}

type targetSchemaController struct {
	service             services.TargetSchemaService
	decisionTreeService services.DecisionTreeService
}

//Function to crete new targetSchema controller
func NewTargetSchemaController() TargetSchemaController {
	return &targetSchemaController{
		service:             services.NewTargetSchemaService(),
		decisionTreeService: services.NewDecisionTreeService(),
	}
}

// GerTargetSchemas godoc
// @Summary Get Target Schemas
// @Schemes
// @Description Get lists of Target Schemas
// @Tags target-schema
// @Produce json
// @Success 200 {array}   models.TargetSchemaSwagger
// @Failure 500 {object}  models.FailedOperation
// @Router /target-schema [get]
func (c *targetSchemaController) GetTargetSchemas(ctx *gin.Context) {
	//Get targetSchemas from db
	targetSchemas, err := c.service.GetTargetSchemas()

	//Handle errors
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.FailedOperation{
			InternalCode: "InternalServerError",
			Message:      err.Error(),
		})
		return
	}

	//Return list of targetShemas
	ctx.JSON(http.StatusOK, targetSchemas)
}

// TestTargetSchemaDecisionTree godoc
// @Summary Test Target Schema Decision Tree
// @Schemes
// @Description Test Target Schema Decision Tree output with Cant and AmountTotal
// @Tags target-schema
// @Accept json
// @Produce json
// @Param DecisionTreeInputs body models.DecisionTreeInputs true "payload"
// @Success 200 {object}  models.SucessfullOperation
// @Failure 400 {object}  models.FailedOperation
// @Failure 500 {object}  models.FailedOperation
// @Router /target-schema/test-tree [post]
func (c *targetSchemaController) TestTargetSchemaDecisionTree(ctx *gin.Context) {
	var decisionTreeInputs models.DecisionTreeInputs

	//Get request body
	err := ctx.ShouldBindJSON(&decisionTreeInputs)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.FailedOperation{
			InternalCode: "BadRequest",
			Message:      err.Error(),
		})
		return
	}

	var targetSchema models.TargetSchema

	//Get latest targetSchema
	err = c.service.FindLatestTargetSchema(&targetSchema)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.FailedOperation{
			InternalCode: "InternalServerError",
			Message:      err.Error(),
		})
		return
	}

	//This function create and test the decision tree with the targetSchema
	target, err := c.decisionTreeService.CreateAndExecuteDecisionTree(
		targetSchema.DesicionTree,
		decisionTreeInputs.Cant,
		decisionTreeInputs.AmountTotal,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.FailedOperation{
			InternalCode: "InternalServerError",
			Message:      err.Error(),
		})
		return
	}

	//Return the target returned by the decision tree
	ctx.JSON(http.StatusOK, models.SucessfullOperation{
		Message: target,
	})
}
