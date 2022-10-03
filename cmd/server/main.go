package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/e-zhydzetski/faraway-tt/internal/app"
)

func errorAwareMain() error {
	cfg := app.DefaultServerConfig()
	if e, present := os.LookupEnv("LISTEN_ADDR"); present {
		cfg.ListenAddr = e
	}
	if e, present := os.LookupEnv("POW_DIFFICULTY"); present {
		d, err := strconv.ParseUint(e, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid POW_DIFFICULTY value: %v", err)
		}
		cfg.Difficulty = d
	}

	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	server, err := app.StartServer(ctx, cfg)
	if err != nil {
		return err
	}
	log.Printf("Server listenting on port: %d", server.Port())

	return server.WaitForShutdown()
}

func main() {
	err := errorAwareMain()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
