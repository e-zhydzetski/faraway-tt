package tcp_test

import (
	"context"
	"fmt"
	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConnection(t *testing.T) {
	ctx := context.Background()

	serverCtx, serverCtxCancel := context.WithTimeout(ctx, 1*time.Second)
	defer serverCtxCancel()

	testBytes := []byte("testBytes")
	testString := "testString"
	testUint64 := uint64(777)

	server, err := tcp.StartServer(serverCtx, ":0", func(ctx context.Context, c *tcp.Connection) error {
		err := c.WriteBytes(testBytes)
		require.NoError(t, err)
		err = c.WriteString(testString)
		require.NoError(t, err)
		err = c.WriteUint64(testUint64)
		require.NoError(t, err)
		return nil
	})
	require.NoError(t, err)

	c, err := tcp.Connect(ctx, fmt.Sprintf("127.0.0.1:%d", server.Port()))
	require.NoError(t, err)

	b, err := c.ReadBytes()
	require.NoError(t, err)
	require.Equal(t, testBytes, b)

	s, err := c.ReadString()
	require.NoError(t, err)
	require.Equal(t, testString, s)

	u, err := c.ReadUint64()
	require.NoError(t, err)
	require.Equal(t, testUint64, u)

	serverCtxCancel()
	_ = c.Close()

	_ = server.WaitForShutdown()
}
