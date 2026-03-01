package main

import (
	"log"
	"net"
)

const (
	response = "OK\n"
	network  = "tcp"
	address  = ":8080"
)

func main() {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Ошибка при Accept: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte(response))
}
