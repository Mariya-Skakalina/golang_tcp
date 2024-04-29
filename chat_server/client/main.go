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
	// net.ResolveTCPAddr используется для преобразования сетевого адреса
	tcpServer, err := net.ResolveTCPAddr(type_serv, host+":"+port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Этот код пытается установить TCP-соединение с удаленным сервером
	conn, err := net.DialTCP(type_serv, nil, tcpServer)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()

	log.Printf("Connected to server %s:%s\n", host, port)

	// Запускаем обработку сообщений от сервера в отдельной горутине
	go handleServerMessages(conn)

	// Отправляем сообщения серверу
	for {
		var message string
		fmt.Print("Enter message: ")
		fmt.Scanln(&message)

		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

// Функция для обработки сообщений от сервера
func handleServerMessages(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		// Читаем сообщение от сервера
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		serverMessage := string(buffer[:n])
		log.Printf("Message from server: %s\n", serverMessage)
	}
}
