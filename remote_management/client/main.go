package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Ошибка подключения к серверу:", err)
	}
	defer conn.Close()

	fmt.Println("Подключен к серверу. Введите команды для выполнения.")
	fmt.Println(`Пример: "ls" или "exit" для выхода.`)

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Ошибка чтения команды:", err)
			return
		}

		_, err = conn.Write([]byte(command))
		if err != nil {
			log.Println("Ошибка отправки команды:", err)
			return
		}

		response, err := serverReader.ReadString('\n')
		if err != nil {
			log.Println("Ошибка чтения ответа от сервера:", err)
			return
		}

		fmt.Print(response)
	}
}
