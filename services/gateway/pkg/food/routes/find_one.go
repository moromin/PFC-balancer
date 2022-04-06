package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/apierror"
	"github.com/moromin/PFC-balancer/services/food/proto"
)

func FindOne(ctx *gin.Context, c proto.FoodServiceClient) {
	name := ctx.Param("name")

	res, err := c.FindOne(context.Background(), &proto.FindOneRequest{
		Name: name,
	})

	if err != nil {
		apierror.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
