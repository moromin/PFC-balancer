package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	auth "github.com/moromin/PFC-balancer/services/auth/proto"
	food "github.com/moromin/PFC-balancer/services/food/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, port int, l *zap.Logger) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	authConn, err := grpc.DialContext(ctx, "localhost:50051", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial to auth server: %w", err)
	}
	if err := auth.RegisterAuthServiceHandlerClient(ctx, mux, auth.NewAuthServiceClient(authConn)); err != nil {
		return fmt.Errorf("failed to regiter auth client: %w", err)
	}

	// TODO: delete menu service
	// menuConn, err := grpc.DialContext(ctx, "localhost:50053", opts...)
	// if err != nil {
	// 	return fmt.Errorf("failed to dial to menu server: %w", err)
	// }
	// if err := menu.RegisterMenuServiceHandlerClient(ctx, mux, menu.NewMenuServiceClient(menuConn)); err != nil {
	// 	return fmt.Errorf("failed to regiter menu client: %w", err)
	// }

	foodConn, err := grpc.DialContext(ctx, "localhost:50055", opts...)
	if err != nil {
		return fmt.Errorf("failed to dial to food server: %w", err)
	}
	if err := food.RegisterFoodServiceHandlerClient(ctx, mux, food.NewFoodServiceClient(foodConn)); err != nil {
		return fmt.Errorf("failed to regiter menu client: %w", err)
	}

	errCh := make(chan error, 1)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("failed to serve rest server: %w", err)
	case <-ctx.Done():
		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown rest server: %w", err)
		}

		if err := <-errCh; err != nil && !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to close rest server: %w", err)
		}

		return nil
	}
}
