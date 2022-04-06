package food

import (
	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/services/gateway/config"
	"github.com/moromin/PFC-balancer/services/gateway/pkg/auth"
	"github.com/moromin/PFC-balancer/services/gateway/pkg/food/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	route := r.Group("/food")
	route.Use(a.AuthRequired)
	route.POST("/", svc.CreateFood)
	route.GET("/", svc.ListFood)
	route.GET("/:name", svc.FindOne)
}

func (svc *ServiceClient) CreateFood(ctx *gin.Context) {
	routes.CreateFood(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) ListFood(ctx *gin.Context) {
	routes.ListFood(ctx, svc.Client)
}
