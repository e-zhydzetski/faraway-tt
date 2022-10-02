package tcp

import (
	"context"
	"log"
	"net"
)

type Handler func(ctx context.Context, c *Connection) error

func ListenAndServe(ctx context.Context, addr string, handler Handler) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		_ = l.Close()
	}()
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go handleConnection(ctx, conn, handler)
	}
}

func handleConnection(ctx context.Context, conn net.Conn, handler Handler) {
	defer conn.Close()
	c := NewConnection(conn)
	err := handler(ctx, c)
	if err != nil {
		log.Println(err)
		_ = c.WriteString("ERROR: " + err.Error())
	}
}
