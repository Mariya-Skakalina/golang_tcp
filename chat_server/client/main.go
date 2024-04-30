package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var (
	host     = "localhost"
	port     = "8000"
	typeServ = "tcp"
)

func main() {
	conn, err := net.Dial(typeServ, host+":"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	go receiveMessages(reader)

	fmt.Print("Введите свое имя: ")
	name, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Fprintf(conn, "%s\n", name)

	for {
		fmt.Print("Введите ваше сообщение: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			break
		}

		fmt.Fprintf(conn, "%s\n", message)
	}
}

func receiveMessages(reader *bufio.Reader) {
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Print(message)
	}
}
