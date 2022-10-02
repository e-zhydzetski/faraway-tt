package tcp

import (
	"context"
	"net"
)

func Connect(_ context.Context, addr string) (*Connection, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return NewConnection(conn), nil
}
