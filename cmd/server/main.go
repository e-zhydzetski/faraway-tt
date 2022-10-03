package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/quoter"

	"github.com/e-zhydzetski/faraway-tt/internal/domain"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/pow"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"
)

func errorAwareMain() error {
	var listenAddr = ":7777"
	if e, present := os.LookupEnv("LISTEN_ADDR"); present {
		listenAddr = e
	}
	var difficulty uint64 = 100
	if e, present := os.LookupEnv("POW_DIFFICULTY"); present {
		d, err := strconv.ParseUint(e, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid POW_DIFFICULTY value: %v", err)
		}
		difficulty = d
	}

	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	powCheckFactory := pow.NewBCryptCheck
	quoterService := quoter.NewQuotableClient()

	handler := domain.NewHandler(powCheckFactory, quoterService, difficulty)

	server, err := tcp.StartServer(ctx, listenAddr, func(ctx context.Context, c *tcp.Connection) error {
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
		os.Exit(1)
	}
}
