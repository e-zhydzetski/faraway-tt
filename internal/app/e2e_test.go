package app

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	server, err := StartServer(ctx, ServerConfig{
		ListenAddr: ":0",
		Difficulty: 1,
	})
	require.NoError(t, err)

	client := NewClient(ClientConfig{
		ServerAddr: fmt.Sprintf("127.0.0.1:%d", server.Port()),
	})

	const clientsCount = 10

	var wg sync.WaitGroup
	wg.Add(clientsCount)
	for i := 0; i < clientsCount; i++ {
		go func() {
			defer wg.Done()
			quote, err := client.RequestForQuote(ctx)
			require.NoError(t, err)
			t.Log(quote)
			require.NotEmpty(t, quote)
		}()
	}
	wg.Wait()

	cancel()
	_ = server.WaitForShutdown()
}
