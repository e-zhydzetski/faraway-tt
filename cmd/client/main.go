package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/pow"
)

func errorAwareMain() error {
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	conn, err := tcp.Connect(ctx, "127.0.0.1:7777")
	if err != nil {
		return err
	}

	powChallenge, err := conn.ReadBytes()
	if err != nil {
		return err
	}

	before := time.Now()
	proof := pow.BCryptProve(powChallenge)
	duration := time.Since(before)
	log.Println("POW check duration:", duration)

	err = conn.WriteUint64(proof)
	if err != nil {
		return err
	}

	quote, err := conn.ReadString()
	if err != nil {
		return err
	}
	log.Println("Quote:", quote)

	return nil
}

func main() {
	err := errorAwareMain()
	if err != nil {
		log.Println(err)
	}
}
