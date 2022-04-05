package food

import (
	"fmt"

	"github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/gateway/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.FoodServiceClient
}

func InitServiceClient(c *config.Config) proto.FoodServiceClient {
	cc, err := grpc.Dial(c.FoodSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewFoodServiceClient(cc)
}
