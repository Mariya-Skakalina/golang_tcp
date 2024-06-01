package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var (
	host     = "localhost"
	port     = "8000"
	typeServ = "tcp"
	// mu       sync.Mutex
)

func init() {
	fmt.Println("Woo")
}

func main() {
	// создание соединения
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

	reader := bufio.NewReader(conn)
	for {
		question_server := "Введите команду (upload, download, exit, all_file): \n"
		if _, err := conn.Write([]byte(question_server)); err != nil {
			log.Println("Ошибка написания команды, ", err)
		}

		command, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("Ошибка чтения команды:", err)
			}
			return
		}
		command = strings.TrimSpace(command)

		switch command {
		case "upload":
			// commander.Upload(conn, reader)
			log.Println("Upload command cloused")
		case "download":
			// download(conn, reader)
			log.Println("Download command cloused")
			continue
		case "all_files":
			all_files(conn, reader)
			continue
		default:
			// log.Println("Неизвестная команда:", command)
			continue
		}
	}
}

func all_files(conn net.Conn, reader *bufio.Reader) {
	c := "Ответ от сервера: Hello all files\n"
	if _, err := conn.Write([]byte(c)); err != nil {
		log.Println("Ошибка на стороне сервера: ", err)
		return
	}
	answer_client, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Ошибка чтения данных от клиента", err)
	}
	fmt.Println(answer_client)

}
