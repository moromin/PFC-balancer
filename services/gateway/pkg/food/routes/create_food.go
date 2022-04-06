package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/moromin/PFC-balancer/apierror"
	"github.com/moromin/PFC-balancer/services/food/proto"
)

type CreateFoodRequestBody struct {
	Name         string  `json:"name" validate:"required"`
	Protein      float64 `json:"protein" validate:"required,gte=0"`
	Fat          float64 `json:"fat" validate:"required,gte=0"`
	Carbohydrate float64 `json:"carbohydrate" validate:"required,gte=0"`
	Category     int64   `json:"category" validate:"required"`
}

func CreateFood(ctx *gin.Context, c proto.FoodServiceClient) {
	body := CreateFoodRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateFood(context.Background(), &proto.CreateFoodRequest{
		Name:         body.Name,
		Protein:      body.Protein,
		Fat:          body.Fat,
		Carbohydrate: body.Carbohydrate,
		Category:     body.Category,
	})

	if err != nil {
		apierror.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
