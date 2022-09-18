package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jpbmdev/payment-api/errorsResponse"
)

func PaginationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ctx.Keys = make(map[string]interface{})
		}
		ctx.Keys["page"] = 1
		ctx.Keys["pageSize"] = 10

		//If page is passed
		if ctx.Query("page") != "" {
			page, err := strconv.Atoi(ctx.Query("page"))
			if err != nil {
				errorsResponse.Error400(ctx, err.Error())
				return
			}
			if page < 0 {
				page = 1
			}
			ctx.Keys["page"] = page
		}

		//If pageSize is passed
		if ctx.Query("pageSize") != "" {
			pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
			if err != nil {
				errorsResponse.Error400(ctx, err.Error())
				return
			}
			if pageSize < 0 {
				pageSize = 1
			}
			ctx.Keys["pageSize"] = pageSize
		}

		ctx.Next()
	}
}
