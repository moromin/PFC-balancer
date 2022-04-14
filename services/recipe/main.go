package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/moromin/PFC-balancer/pkg/logger"
	"github.com/moromin/PFC-balancer/services/recipe/grpc"
)

var port = flag.Int("p", 4000, "gRPC server network port")
var dbAddr = flag.String("dbAddr", "localhost:50000", "DB service address")
var authAddr = flag.String("authAddr", "localhost:50051", "Auth service address")
var userAddr = flag.String("userAddr", "localhost:50052", "User service address")
var foodAddr = flag.String("foodAddr", "localhost:50054", "Food service address")

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	l, err := logger.New()
	if err != nil {
		_, ferr := fmt.Fprintf(os.Stderr, "failed to create logger: %s", err)
		if ferr != nil {
			log.Fatalf("failed to write log: %s, original error: %s", err, ferr)
		}
		return 1
	}

	flag.Parse()
	cfg := &grpc.ServerConfig{
		Port:     *port,
		DBAddr:   *dbAddr,
		AuthAddr: *authAddr,
		UserAddr: *userAddr,
		FoodAddr: *foodAddr,
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- grpc.RunServer(ctx, cfg, l)
	}()

	select {
	case err := <-errCh:
		log.Println(err)
		return 1
	case <-ctx.Done():
		log.Println("shutting down...")
		return 0
	}
}
