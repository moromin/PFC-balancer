package grpc

import (
	"context"

	pkggrpc "github.com/moromin/PFC-balancer/pkg/grpc"
	food "github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/menu/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, l *zap.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	conn, err := grpc.DialContext(ctx, "localhost:50055", opts...)
	if err != nil {
		return err
	}

	foodClient := food.NewFoodServiceClient(conn)

	svc := &server{
		foodClient: foodClient,
	}

	return pkggrpc.NewServer(port, func(s *grpc.Server) {
		proto.RegisterMenuServiceServer(s, svc)
	}).Start(ctx)
}
