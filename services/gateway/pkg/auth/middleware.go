package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/services/auth/proto"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("authorization header does not exist"))
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("incorrect number of token"))
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &proto.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("your token is invalid"))
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
