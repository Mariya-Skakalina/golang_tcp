package main

import (
	"log"
	"net"
	"os"
)

var (
	host      = "localhost"
	port      = "8000"
	type_serv = "tcp"
)

// Мапа для хранения соединений клиентов
var connsClient = make(map[net.Conn]string)

func main() {
	// Устанавливаем соединение
	listen, err := net.Listen(type_serv, host+":"+port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer listen.Close()

	log.Printf("Server started at %s:%s\n", host, port)

	// Запускаем обработку подключений в отдельной горутине
	go handleConnections(listen)

	// Бесконечный цикл для ожидания сигнала завершения программы
	select {}
}

// Функция для обработки подключений клиентов
func handleConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Printf("Client connected from %s\n", conn.RemoteAddr())

		// Добавляем соединение клиента в мапу
		connsClient[conn] = conn.RemoteAddr().String()

		// Запускаем обработку сообщений от клиента в отдельной горутине
		go handleClientMessages(conn)
	}
}

// Функция для обработки сообщений от клиента
func handleClientMessages(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		// Читаем сообщение от клиента
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		clientMessage := string(buffer[:n])
		log.Printf("Message from %s: %s\n", conn.RemoteAddr(), clientMessage)

		// Отправляем сообщение от клиента всем остальным клиентам
		for c, _ := range connsClient {
			if c != conn {
				_, err := c.Write([]byte(clientMessage))
				if err != nil {
					log.Println("Write error:", err)
				}
			}
		}
	}
}
