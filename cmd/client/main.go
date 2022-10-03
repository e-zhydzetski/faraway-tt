package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	farawaytt "github.com/e-zhydzetski/faraway-tt/pkg"
)

func errorAwareMain() error {
	cfg := farawaytt.DefaultClientConfig()
	flag.StringVar(&cfg.ServerAddr, "server-addr", cfg.ServerAddr, "server address")
	flag.Parse()

	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	client := farawaytt.NewClient(cfg)
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
