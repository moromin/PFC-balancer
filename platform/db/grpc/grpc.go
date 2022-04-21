package grpc

import (
	"context"

	pkggrpc "github.com/moromin/PFC-balancer/pkg/grpc"
	"github.com/moromin/PFC-balancer/platform/db/config"
	"github.com/moromin/PFC-balancer/platform/db/db"
	"github.com/moromin/PFC-balancer/platform/db/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, l *zap.Logger) error {
	c := config.LoadConfig()
	db, err := db.New(&c)
	if err != nil {
		return err
	}

	svc := &server{
		db: db,
	}

	return pkggrpc.NewServer(port, l, func(s *grpc.Server) {
		proto.RegisterDBServiceServer(s, svc)
	}).Start(ctx)
}
