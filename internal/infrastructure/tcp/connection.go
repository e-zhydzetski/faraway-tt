package tcp

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
)

func NewConnection(conn baseConn) *Connection {
	return &Connection{conn: conn}
}

type Connection struct {
	conn baseConn
}

func (c *Connection) WriteUint64(data uint64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, data)
	_, err := c.conn.Write(b)
	return err
}

func (c *Connection) ReadUint64(ctx context.Context) (uint64, error) {
	go func() {
		<-ctx.Done()
		_ = c.Close()
	}()
	b := make([]byte, 8)
	n, err := c.conn.Read(b)
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, errors.New("invalid size read from connection, expected 8 bytes")
	}
	return binary.LittleEndian.Uint64(b), nil
}

func (c *Connection) WriteBytes(data []byte) error {
	err := c.WriteUint64(uint64(len(data)))
	if err != nil {
		return err
	}
	_, err = c.conn.Write(data)
	return err
}

const maxBufferSize = 4096

func (c *Connection) ReadBytes(ctx context.Context) ([]byte, error) {
	go func() {
		<-ctx.Done()
		_ = c.Close()
	}()
	l, err := c.ReadUint64(ctx)
	if err != nil {
		return nil, err
	}
	if l > maxBufferSize {
		return nil, fmt.Errorf("read bytes error: max buffer size exceeded: max %d, cur %d", maxBufferSize, l)
	}
	b := make([]byte, l)
	n, err := c.conn.Read(b)
	if err != nil {
		return nil, err
	}
	if uint64(n) != l {
		return nil, fmt.Errorf("invalid size read from connection, expected %d bytes", l)
	}
	return b, nil
}

func (c *Connection) WriteString(data string) error {
	return c.WriteBytes([]byte(data))
}

func (c *Connection) ReadString(ctx context.Context) (string, error) {
	b, err := c.ReadBytes(ctx)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (c *Connection) Close() error {
	return c.conn.Close()
}
