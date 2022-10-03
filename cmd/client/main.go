package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/e-zhydzetski/faraway-tt/internal/app"
)

func errorAwareMain() error {
	cfg := app.DefaultClientConfig()
	flag.StringVar(&cfg.ServerAddr, "server-addr", cfg.ServerAddr, "server address")
	flag.Parse()

	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	client := app.NewClient(cfg)
	quote, err := client.RequestForQuote(ctx)
	if err != nil {
		return err
	}
	log.Println("Quote:", quote)

	return nil
}

func main() {
	err := errorAwareMain()
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}
}
