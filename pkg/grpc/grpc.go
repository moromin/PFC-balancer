package grpc

import (
	"context"
	"fmt"
	"net"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	logger "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	port   int
}

func NewServer(port int, register func(*grpc.Server)) *Server {
	zl, _ := zap.NewProduction()
	logger.ReplaceGrpcLogger(zl)

	intersections := []grpc.UnaryServerInterceptor{
		ctxtags.UnaryServerInterceptor(ctxtags.WithFieldExtractor(ctxtags.CodeGenRequestFieldExtractor)),
		logger.UnaryServerInterceptor(zl),
	}

	opts := []grpc.ServerOption{
		middleware.WithUnaryServerChain(intersections...),
	}

	server := grpc.NewServer(opts...)

	register(server)

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
