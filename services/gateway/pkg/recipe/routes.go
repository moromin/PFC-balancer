package recipe

import (
	"github.com/gin-gonic/gin"
	"github.com/moromin/PFC-balancer/services/gateway/config"
	"github.com/moromin/PFC-balancer/services/gateway/pkg/auth"
	"github.com/moromin/PFC-balancer/services/gateway/pkg/recipe/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	route := r.Group("/recipe")
	route.Use(a.AuthRequired)
	route.POST("/", svc.CreateRecipe)
	route.GET("/:id", svc.ReadRecipe)
}

func (svc *ServiceClient) CreateRecipe(ctx *gin.Context) {
	routes.CreateRecipe(ctx, svc.Client)
}

func (svc *ServiceClient) ReadRecipe(ctx *gin.Context) {
	routes.ReadRecipe(ctx, svc.Client)
}
