package auth

import (
	"fmt"

	"github.com/moromin/PFC-balancer/services/auth/proto"
	"github.com/moromin/PFC-balancer/services/gateway/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.AuthServiceClient
}

func InitServiceClient(c *config.Config) proto.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewAuthServiceClient(cc)
}
