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
	route.GET("/", svc.ListFoods)
	route.GET("/:name", svc.FindOne)
	route.GET("/search/:name", svc.SearchFoods)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) ListFoods(ctx *gin.Context) {
	routes.ListFoods(ctx, svc.Client)
}

func (svc *ServiceClient) SearchFoods(ctx *gin.Context) {
	routes.SearchFoods(ctx, svc.Client)
}
