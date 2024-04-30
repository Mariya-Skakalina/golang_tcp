package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

var (
	host     = "localhost"
	port     = "8000"
	typeServ = "tcp"
)

// Определяем структуру клиента
type Client struct {
	Name string
	Conn net.Conn
}

// Список клиентов
var clients []Client
var mu sync.Mutex

func main() {
	// Установка соединения
	listener, err := net.Listen(typeServ, host+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Printf("Server started on %s:%s\n", host, port)

	// Подключение
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Чтение пользовательских данных
	reader := bufio.NewReader(conn)
	fmt.Fprintf(conn, "Сообщение отправил пользователь по имени, ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	client := Client{Name: name, Conn: conn}

	mu.Lock()
	clients = append(clients, client)
	mu.Unlock()

	log.Printf("%s connected\n", name)

	// Чтение сообщений
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}

		message = strings.TrimSpace(message)

		if message == "exit" {
			break
		}

		broadcastMessage(name, message)
	}

	removeClient(client)
	log.Printf("%s disconnected\n", name)
}

// Рассылка сообщений всем клиентам
func broadcastMessage(senderName, message string) {
	mu.Lock()
	defer mu.Unlock()

	for _, client := range clients {
		if client.Name != senderName {
			fmt.Fprintf(client.Conn, "%s: %s\n", senderName, message)
		}
	}
}

// Удаление вышедшего клиента
func removeClient(client Client) {
	mu.Lock()
	defer mu.Unlock()

	for i, c := range clients {
		if c == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}
