package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	auth "github.com/moromin/PFC-balancer/services/auth/proto"
	food "github.com/moromin/PFC-balancer/services/food/proto"
	recipe "github.com/moromin/PFC-balancer/services/recipe/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type ServerConfig struct {
	Port       int
	AuthAddr   string
	RecipeAddr string
	FoodAddr   string
}

func RunServer(ctx context.Context, cfg *ServerConfig, l *zap.Logger) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	authConn, err := grpc.DialContext(ctx, cfg.AuthAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial to auth server: %w", err)
	}
	if err := auth.RegisterAuthServiceHandlerClient(ctx, mux, auth.NewAuthServiceClient(authConn)); err != nil {
		return fmt.Errorf("failed to regiter auth client: %w", err)
	}

	foodConn, err := grpc.DialContext(ctx, cfg.FoodAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial to food server: %w", err)
	}
	if err := food.RegisterFoodServiceHandlerClient(ctx, mux, food.NewFoodServiceClient(foodConn)); err != nil {
		return fmt.Errorf("failed to regiter food client: %w", err)
	}

	recipeConn, err := grpc.DialContext(ctx, cfg.RecipeAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial to recipe server: %w", err)
	}
	if err := recipe.RegisterRecipeServiceHandlerClient(ctx, mux, recipe.NewRecipeServiceClient(recipeConn)); err != nil {
		return fmt.Errorf("failed to regiter recipe client: %w", err)
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
