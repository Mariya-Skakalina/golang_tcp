package main

import (
	"fmt"
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
		fmt.Println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	// Этот код пытается установить TCP-соединение с удаленным сервером
	for {
		conn, err := net.DialTCP(type_serv, nil, tcpServer)
		defer conn.Close()
		if err != nil {
			fmt.Println("Dial failed:", err.Error())
			os.Exit(1)
		}
		handlerRequest(conn)
	}
}

func handlerRequest(conn net.Conn) {
	// Читает сообщение пользователя
	var message string
	fmt.Scan(&message)

	// Отправляет сообщение на сервер
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// Этот код читает данные из соединения TCP и сохраняет их в буфер received.
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		fmt.Println("Read data failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println(string(received))
}
