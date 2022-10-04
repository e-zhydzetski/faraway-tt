package tcp_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"
)

func WithTimeout(timeout time.Duration, f func(t *testing.T)) func(t *testing.T) {
	return func(t *testing.T) {
		complete := make(chan struct{})
		go func() {
			f(t)
			close(complete)
		}()

		select {
		case <-time.After(timeout):
			t.Fatal("timeout")
		case <-complete:
		}
	}
}

func TestServerCtxCancel(t *testing.T) {
	WithTimeout(time.Second, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		server, err := tcp.StartServer(ctx, ":0", func(ctx context.Context, c *tcp.Connection) error {
			_, err := c.ReadUint64() // wait forever
			require.ErrorContains(t, err, "use of closed network connection")
			return err
		})
		require.NoError(t, err)

		_, err = tcp.Connect(ctx, fmt.Sprintf("127.0.0.1:%d", server.Port()))
		require.NoError(t, err)
		cancel()

		_ = server.WaitForShutdown()
	})(t)
}

func TestServerConnectionClose(t *testing.T) {
	WithTimeout(time.Second, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		done := make(chan struct{})
		server, err := tcp.StartServer(ctx, ":0", func(ctx context.Context, c *tcp.Connection) error {
			err := c.Close()
			<-ctx.Done()
			close(done)
			return err
		})
		require.NoError(t, err)

		_, err = tcp.Connect(ctx, fmt.Sprintf("127.0.0.1:%d", server.Port()))
		require.NoError(t, err)

		<-done
		cancel()

		_ = server.WaitForShutdown()
	})(t)
}
