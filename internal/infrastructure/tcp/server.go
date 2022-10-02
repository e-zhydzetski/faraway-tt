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
					log.Printf("accept tmp error: %v", err)
					continue
				}
				return err
			}
			g.Go(func() error { // start inside errgroup to wait all in-fly connections after shutdown
				defer conn.Close()
				c := NewConnection(conn)
				// TODO maybe use another context for in-fly requests
				err := handler(ctx, c)
				if err != nil {
					log.Println(err)
					_ = c.WriteString("ERROR: " + err.Error())
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
