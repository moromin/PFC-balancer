package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/apierror"
	"github.com/moromin/PFC-balancer/services/food/proto"
)

func FindOne(ctx *gin.Context, c proto.FoodServiceClient) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("failed to convert id"))
	}

	res, err := c.FindOne(context.Background(), &proto.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		apierror.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
