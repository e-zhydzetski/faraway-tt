package main

import (
	"log"
	"net"
	"time"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/tcp"

	"github.com/e-zhydzetski/faraway-tt/internal/infrastructure/pow"
)

func errorAwareMain() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7777")
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	connection := tcp.NewConnection(conn)

	powChallenge, err := connection.ReadBytes()
	if err != nil {
		return err
	}

	before := time.Now()
	proof := pow.BCryptProve(powChallenge)
	duration := time.Since(before)
	log.Println("POW check duration:", duration)

	err = connection.WriteUint64(proof)
	if err != nil {
		return err
	}

	quote, err := connection.ReadString()
	if err != nil {
		return err
	}
	log.Println("Quote:", quote)

	return nil
}

func main() {
	err := errorAwareMain()
	if err != nil {
		log.Println(err)
	}
}
