package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	network = "tcp"
	address = "localhost:8080"
)

func main() {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if scanner.Text() == "OK" {
			fmt.Println("Сервер работает корректно")
		}
	}
}
