package grpc

import (
	"context"

	pkggrpc "github.com/moromin/PFC-balancer/pkg/grpc"
	db "github.com/moromin/PFC-balancer/platform/db/proto"
	"github.com/moromin/PFC-balancer/services/food/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, l *zap.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	conn, err := grpc.DialContext(ctx, "localhost:5000", opts...)
	if err != nil {
		return err
	}

	dbClient := db.NewDBServiceClient(conn)

	svc := &server{
		dbClient: dbClient,
	}

	return pkggrpc.NewServer(port, func(s *grpc.Server) {
		proto.RegisterFoodServiceServer(s, svc)
	}).Start(ctx)
}
