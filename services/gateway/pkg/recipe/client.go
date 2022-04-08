package recipe

import (
	"log"

	"github.com/moromin/PFC-balancer/services/gateway/config"
	"github.com/moromin/PFC-balancer/services/recipe/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.RecipeServiceClient
}

func InitServiceClient(c *config.Config) proto.RecipeServiceClient {
	cc, err := grpc.Dial(c.RecipeSvcUrl, grpc.WithInsecure())

	if err != nil {
		log.Println("Could not connect:", err)
	}

	return proto.NewRecipeServiceClient(cc)
}
