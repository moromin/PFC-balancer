package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/moromin/PFC-balancer/pkg/logger"
	"github.com/moromin/PFC-balancer/services/gateway/rest"
)

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

	errCh := make(chan error, 1)
	go func() {
		errCh <- rest.RunServer(ctx, 4000, l)
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
