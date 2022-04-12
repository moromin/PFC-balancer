package grpc

import (
	"context"

	pkggrpc "github.com/moromin/PFC-balancer/pkg/grpc"
	food "github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/menu/proto"
	recipe "github.com/moromin/PFC-balancer/services/recipe/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, l *zap.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	foodConn, err := grpc.DialContext(ctx, "localhost:50055", opts...)
	if err != nil {
		return err
	}
	foodClient := food.NewFoodServiceClient(foodConn)

	recipeConn, err := grpc.DialContext(ctx, "localhost:50054", opts...)
	if err != nil {
		return err
	}
	recipeClient := recipe.NewRecipeServiceClient(recipeConn)

	svc := &server{
		foodClient:   foodClient,
		recipeClient: recipeClient,
	}

	return pkggrpc.NewServer(port, func(s *grpc.Server) {
		proto.RegisterMenuServiceServer(s, svc)
	}).Start(ctx)
}
