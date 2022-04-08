package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/apierror"
	"github.com/moromin/PFC-balancer/services/recipe/proto"
)

func CreateRecipe(ctx *gin.Context, c proto.RecipeServiceClient) {
	req := proto.CreateRecipeRequest{}

	if err := ctx.BindJSON(&req.Data); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")
	req.UserId = userId.(int64)

	res, err := c.CreateRecipe(context.Background(), &req)
	if err != nil {
		apierror.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
