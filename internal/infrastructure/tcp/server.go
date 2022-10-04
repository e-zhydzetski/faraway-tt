package tcp

import (
	"context"
	"log"
	"net"

	"golang.org/x/sync/errgroup"
)

type Handler func(ctx context.Context, c *Connection) error

func StartServer(ctx context.Context, addr string, handler Handler) (*Server, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		<-ctx.Done()
		_ = ln.Close()
		return ctx.Err()
	})
	g.Go(func() error {
		for {
			conn, err := ln.Accept()
			if err != nil {
				if ne, ok := err.(net.Error); ok && ne.Timeout() {
					log.Printf("Accept tmp error: %v", err)
					continue
				}
				return err
			}
			g.Go(func() error { // start inside errgroup to wait all in-fly connections after shutdown
				// TODO maybe use another base context for in-fly requests
				conn, connCtx, cancel := wrapConnectionWithContext(conn, ctx)
				defer cancel()
				err := handler(connCtx, NewConnection(conn))
				if err != nil {
					log.Println("Connection handling error:", err)
				}
				return nil // always return nil to prevent server stop because of connection handling error
			})
		}
	})
	return &Server{
		addr:     ln.Addr().(*net.TCPAddr),
		errGroup: g,
	}, nil
}

type Server struct {
	addr     *net.TCPAddr
	errGroup *errgroup.Group
}

func (s *Server) Port() int {
	return s.addr.Port
}

func (s *Server) WaitForShutdown() error {
	return s.errGroup.Wait()
}
