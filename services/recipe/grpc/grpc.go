package grpc

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	pkggrpc "github.com/moromin/PFC-balancer/pkg/grpc"
	db "github.com/moromin/PFC-balancer/platform/db/proto"
	auth "github.com/moromin/PFC-balancer/services/auth/proto"
	food "github.com/moromin/PFC-balancer/services/food/proto"
	"github.com/moromin/PFC-balancer/services/recipe/proto"
	user "github.com/moromin/PFC-balancer/services/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ServerConfig struct {
	Port     int
	DBAddr   string
	AuthAddr string
	UserAddr string
	FoodAddr string
}

func RunServer(ctx context.Context, cfg *ServerConfig, l *zap.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	dbConn, err := grpc.DialContext(ctx, cfg.DBAddr, opts...)
	if err != nil {
		return err
	}
	dbClient := db.NewDBServiceClient(dbConn)

	authConn, err := grpc.DialContext(ctx, cfg.AuthAddr, opts...)
	if err != nil {
		return err
	}
	authClient := auth.NewAuthServiceClient(authConn)

	userConn, err := grpc.DialContext(ctx, cfg.UserAddr, opts...)
	if err != nil {
		return err
	}
	userClient := user.NewUserServiceClient(userConn)

	foodConn, err := grpc.DialContext(ctx, cfg.FoodAddr, opts...)
	if err != nil {
		return err
	}
	foodClient := food.NewFoodServiceClient(foodConn)

	svc := &server{
		dbClient:   dbClient,
		userClient: userClient,
		foodClient: foodClient,
		authClient: authClient,
	}

	return pkggrpc.NewServer(cfg.Port, func(s *grpc.Server) {
		proto.RegisterRecipeServiceServer(s, svc)
	}, grpc_auth.UnaryServerInterceptor(svc.Authenticate)).Start(ctx)
}
