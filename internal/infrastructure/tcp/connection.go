package tcp

import (
	"context"
	"encoding/binary"
	"net"
)

func NewConnection(conn net.Conn) *Connection {
	return &Connection{conn: conn}
}

type Connection struct {
	conn net.Conn
}

func (c *Connection) POWVerification(ctx context.Context, input []byte) (uint64, error) {
	_, err := c.conn.Write(input)
	if err != nil {
		return 0, err
	}
	proofBytes := make([]byte, 8) // uint64
	_, err = c.conn.Read(proofBytes)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(proofBytes), nil
}

func (c *Connection) SendQuote(ctx context.Context, quote string) error {
	_, err := c.conn.Write([]byte(quote))
	return err
}

func (c *Connection) ReportError(err error) {
	_, _ = c.conn.Write([]byte(err.Error()))
}
