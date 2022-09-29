package main

import (
	"encoding/binary"
	"log"
	"net"
	"time"

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

	buff := make([]byte, 1024)

	n, err := conn.Read(buff)
	if err != nil {
		return err
	}
	powInput := buff[:n]

	before := time.Now()
	answer := pow.BCryptSolve(powInput)
	duration := time.Since(before)
	log.Println("POW check duration:", duration)

	binary.BigEndian.PutUint64(buff, answer)
	_, err = conn.Write(buff[:8])
	if err != nil {
		return err
	}

	n, err = conn.Read(buff)
	if err != nil {
		return err
	}
	quote := buff[:n]
	log.Println("Quote:", string(quote))
	return nil
}

func main() {
	err := errorAwareMain()
	if err != nil {
		log.Println(err)
	}
}
