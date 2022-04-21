package grpc

import (
	"context"
	"fmt"
	"net"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	ctxtags "github.com/moromin/PFC-balancer/pkg/grpc/context"
	logger "github.com/moromin/PFC-balancer/pkg/grpc/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	server *grpc.Server
	port   int
}

func NewServer(port int, zl *zap.Logger, register func(*grpc.Server), iss ...grpc.UnaryServerInterceptor) *Server {
	intersections := []grpc.UnaryServerInterceptor{
		ctxtags.UnaryServerInterceptor(),
		logger.UnaryServerInterceptor(zl),
	}
	intersections = append(intersections, iss...)

	opts := []grpc.ServerOption{
		middleware.WithUnaryServerChain(intersections...),
	}

	server := grpc.NewServer(opts...)

	register(server)
	reflection.Register(server)

	return &Server{
		server: server,
		port:   port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", s.port, err)
	}

	errCh := make(chan error, 1)
	go func() {
		if err := s.server.Serve(lis); err != nil {
			errCh <- err
		}
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("failed to serve %w", err)
		}
		return nil

	case <-ctx.Done():
		s.server.GracefulStop()
		return <-errCh
	}
}
