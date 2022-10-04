package tcp

import (
	"context"
	"net"
)

func Connect(ctx context.Context, addr string) (*Connection, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	var conn baseConn
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	conn, _, _ = wrapConnectionWithContext(conn, ctx)
	return NewConnection(conn), nil
}
