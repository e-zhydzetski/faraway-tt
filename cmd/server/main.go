package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/quoter"

	"github.com/e-zhydzetski/faraway-tt/internal/domain"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/pow"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"
)

func errorAwareMain() error {
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	powCheckFactory := pow.NewBCryptCheck
	quoterService := quoter.NewQuotableClient()

	handler := domain.NewHandler(powCheckFactory, quoterService)

	server, err := tcp.StartServer(ctx, ":7777", func(ctx context.Context, c *tcp.Connection) error {
		return handler.Handle(ctx, c)
	})
	if err != nil {
		return err
	}
	log.Printf("TCP server listenting on port: %d", server.Port())

	return server.WaitForShutdown()
}

func main() {
	err := errorAwareMain()
	if err != nil {
		log.Println(err)
	}
}
