package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/apierror"
	"github.com/moromin/PFC-balancer/services/food/proto"
)

func ListFood(ctx *gin.Context, c proto.FoodServiceClient) {
	res, err := c.ListFood(context.Background(), &proto.ListFoodRequest{})

	if err != nil {
		apierror.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
