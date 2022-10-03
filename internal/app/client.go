package app

import (
	"context"
	"log"
	"time"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/pow"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"
)

type ClientConfig struct {
	ServerAddr string
	DebugLog   bool
}

func DefaultClientConfig() ClientConfig {
	return ClientConfig{
		ServerAddr: "127.0.0.1:7777",
		DebugLog:   true,
	}
}

type Client interface {
	RequestForQuote(ctx context.Context) (string, error)
}

func NewClient(cfg ClientConfig) Client {
	return &client{
		cfg: cfg,
	}
}

type client struct {
	cfg ClientConfig
}

func (c *client) RequestForQuote(ctx context.Context) (string, error) {
	conn, err := tcp.Connect(ctx, c.cfg.ServerAddr)
	if err != nil {
		return "", err
	}

	powChallenge, err := conn.ReadBytes()
	if err != nil {
		return "", err
	}

	proof := func() uint64 {
		if c.cfg.DebugLog {
			defer func(start time.Time) {
				log.Println("POW check duration:", time.Since(start))
			}(time.Now())
		}
		return pow.BCryptProve(powChallenge)
	}()

	err = conn.WriteUint64(proof)
	if err != nil {
		return "", err
	}

	quote, err := conn.ReadString()
	if err != nil {
		return "", err
	}
	return quote, nil
}
