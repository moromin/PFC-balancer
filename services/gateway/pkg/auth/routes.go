package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/services/gateway/config"
	"github.com/moromin/PFC-balancer/services/gateway/pkg/auth/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	route := r.Group("/auth")
	route.POST("/register", svc.Register)
	route.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
