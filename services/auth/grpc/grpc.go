package grpc

import (
	"context"

	pkggrpc "github.com/moromin/PFC-balancer/pkg/grpc"
	"github.com/moromin/PFC-balancer/services/auth/proto"
	"github.com/moromin/PFC-balancer/services/auth/utils"
	user "github.com/moromin/PFC-balancer/services/user/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ServerConfig struct {
	Port     int
	UserAddr string
	Logger   *zap.Logger
}

func RunServer(ctx context.Context, cfg *ServerConfig, l *zap.Logger) error {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	conn, err := grpc.DialContext(ctx, cfg.UserAddr, opts...)
	if err != nil {
		return err
	}

	userClient := user.NewUserServiceClient(conn)
	jwt := utils.JwtWrapper{
		Issuer:          "go-grpc-auth-service",
		ExpirationHours: 24,
	}

	svc := &server{
		userClient: userClient,
		jwt:        jwt,
	}

	return pkggrpc.NewServer(cfg.Port, cfg.Logger, func(s *grpc.Server) {
		proto.RegisterAuthServiceServer(s, svc)
	}).Start(ctx)
}
