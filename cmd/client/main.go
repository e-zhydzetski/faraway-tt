package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/pow"
)

func errorAwareMain() error {
	var serverAddr string
	flag.StringVar(&serverAddr, "server-addr", "127.0.0.1:7777", "server address")
	flag.Parse()

	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	conn, err := tcp.Connect(ctx, serverAddr)
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
		log.Println("ERROR:", err)
		os.Exit(1)
	}
}
