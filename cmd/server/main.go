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

func main() {
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	powCheckFactory := pow.NewBCryptCheck
	quoter := quoter.NewQuotableClient()

	handler := domain.NewHandler(powCheckFactory, quoter)

	err := tcp.ListenAndServe(ctx, ":7777", handler)
	log.Println(err)
}
