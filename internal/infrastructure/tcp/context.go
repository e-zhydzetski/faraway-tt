package tcp

import (
	"context"
)

//nolint:revive // ctx here should be the second parameter
func wrapConnectionWithContext(conn baseConn, ctx context.Context) (baseConn, context.Context, context.CancelFunc) {
	ctx, ctxCancel := context.WithCancel(ctx)

	c := &ContextAwareConnection{
		conn:      conn,
		ctxCancel: ctxCancel,
	}

	go func() {
		<-ctx.Done()
		_ = c.Close()
	}()

	return c, ctx, ctxCancel
}

type ContextAwareConnection struct {
	conn      baseConn
	ctxCancel context.CancelFunc
}

func (c *ContextAwareConnection) Close() error {
	c.ctxCancel()
	return c.conn.Close()
}

func (c *ContextAwareConnection) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}

func (c *ContextAwareConnection) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}
