package tcp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

func NewConnection(conn net.Conn) *Connection {
	return &Connection{conn: conn}
}

type Connection struct {
	conn net.Conn
}

// Transport level methods

func (c *Connection) WriteUint64(data uint64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, data)
	_, err := c.conn.Write(b)
	return err
}

func (c *Connection) ReadUint64() (uint64, error) {
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

func (c *Connection) ReadBytes() ([]byte, error) {
	l, err := c.ReadUint64()
	if err != nil {
		return nil, err
	}
	// TODO check max len and use buffer
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

func (c *Connection) ReadString() (string, error) {
	b, err := c.ReadBytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
