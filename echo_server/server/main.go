package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

var (
	host      = "localhost"
	port      = "8000"
	type_serv = "tcp"
)

func main() {
	// Усанавливем соединение
	listen, err := net.Listen(type_serv, host+":"+port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer listen.Close()
	// Запускаем его в бесконечном цикле
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		// Запуск нескольких клиентов
		go handlerRequest(conn)
	}

}

func handlerRequest(conn net.Conn) {
	defer conn.Close()

	log.Printf("Client connected from %s\n", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Read error:", err)
		return
	}
	clientMessage := string(buffer[:n])
	responseStr := fmt.Sprintf("Your message is: %v.", clientMessage)
	_, err = conn.Write([]byte(responseStr))
	if err != nil {
		log.Println("Write error:", err)
	}
}
