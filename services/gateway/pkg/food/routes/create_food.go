package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/services/food/proto"
)

type CreateFoodRequestBody struct {
	Name         string  `json:"name"`
	Protein      float64 `json:"protein"`
	Fat          float64 `json:"fat"`
	Carbohydrate float64 `json:"carbohydrate"`
	Category     int64   `json:"category"`
}

func CreateFood(ctx *gin.Context, c proto.FoodServiceClient) {
	body := CreateFoodRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
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
		ctx.AbortWithError(int(res.Status), err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
