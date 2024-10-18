package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/alxrusinov/diploma/internal/app"
)

func main() {
	app := app.NewApp()

	ctx, cancel := context.WithCancel(context.Background())

	signalCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	errChan := app.Run(ctx)

	select {
	case err := <-errChan:
		log.Fatal("server has been crashed", err)
	case <-signalCtx.Done():
		cancel()

	}

}
