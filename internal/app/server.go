package app

import (
	"context"

	"github.com/e-zhydzetski/faraway-tt/internal/domain"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/pow"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/quoter"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"
)

type ServerConfig struct {
	ListenAddr string
	Difficulty uint64
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		ListenAddr: ":7777",
		Difficulty: 100,
	}
}

type Server interface {
	Port() int
	WaitForShutdown() error
}

func StartServer(ctx context.Context, cfg ServerConfig) (Server, error) {
	powCheckFactory := pow.NewBCryptCheck
	quoterService := quoter.NewQuotableClient()

	service := domain.NewService(powCheckFactory, quoterService, cfg.Difficulty)

	return tcp.StartServer(ctx, cfg.ListenAddr, func(ctx context.Context, c *tcp.Connection) error {
		return service.ServeClient(ctx, c)
	})
}
