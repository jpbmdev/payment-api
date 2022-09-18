package errorsResponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/models"
)

func Error400(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, models.FailedOperation{
		InternalCode: "BadRequest",
		Message:      message,
	})
}

func Error404(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, models.FailedOperation{
		InternalCode: "NotFound",
		Message:      message,
	})
}

func Error500(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, models.FailedOperation{
		InternalCode: "InternalServerError",
		Message:      message,
	})
}
