package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/apierror"
	"github.com/moromin/PFC-balancer/services/recipe/proto"
)

func ReadRecipe(ctx *gin.Context, c proto.RecipeServiceClient) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid recipe id"))
		return
	}

	userId, _ := ctx.Get("userId")
	res, err := c.ReadRecipe(context.Background(), &proto.ReadRecipeRequest{
		Id:     int64(id),
		UserId: userId.(int64),
	})

	if err != nil {
		apierror.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
