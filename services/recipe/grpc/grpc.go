package grpc

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	pkggrpc "github.com/moromin/PFC-balancer/pkg/grpc"
	db "github.com/moromin/PFC-balancer/platform/db/proto"
	food "github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/recipe/proto"
	user "github.com/moromin/PFC-balancer/services/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, l *zap.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	dbConn, err := grpc.DialContext(ctx, "localhost:5000", opts...)
	if err != nil {
		return err
	}
	dbClient := db.NewDBServiceClient(dbConn)

	userConn, err := grpc.DialContext(ctx, "localhost:50052", opts...)
	if err != nil {
		return err
	}
	userClient := user.NewUserServiceClient(userConn)

	foodConn, err := grpc.DialContext(ctx, "localhost:50055", opts...)
	if err != nil {
		return err
	}
	foodClient := food.NewFoodServiceClient(foodConn)

	svc := &server{
		dbClient:   dbClient,
		userClient: userClient,
		foodClient: foodClient,
	}

	return pkggrpc.NewServer(port, func(s *grpc.Server) {
		proto.RegisterRecipeServiceServer(s, svc)
	}, grpc_auth.UnaryServerInterceptor(svc.Authenticate)).Start(ctx)
}
