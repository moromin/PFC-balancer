package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/moromin/PFC-balancer/apierror"
	"github.com/moromin/PFC-balancer/services/auth/proto"
)

type RegisterRequestBody struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c proto.AuthServiceClient) {
	body := RegisterRequestBody{}

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

	res, err := c.Register(context.Background(), &proto.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		apierror.AbortWithError(ctx, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
