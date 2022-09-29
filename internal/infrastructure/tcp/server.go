package tcp

import (
	"context"
	"log"
	"net"

	"github.com/e-zhydzetski/faraway-tt/internal/domain"
)

type Handler func(ctx context.Context, c *Connection)

func ListenAndServe(ctx context.Context, addr string, handler domain.Handler) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		l.Close()
	}()
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go handleConnection(ctx, conn, handler)
	}
}

func handleConnection(ctx context.Context, conn net.Conn, handler domain.Handler) {
	defer conn.Close()
	c := NewConnection(conn)
	err := handler.Handle(ctx, c)
	if err != nil {
		log.Println(err)
		c.ReportError(err)
	}
}
