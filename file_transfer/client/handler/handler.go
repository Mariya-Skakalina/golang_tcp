package handler

import (
	"bufio"
	"file_transfer/client/commands"
	"fmt"
	"log"
	"net"
)

func HandlerRequest(conn net.Conn) {

	request := bufio.NewReader(conn)
	responce, err := request.ReadString('\n')
	if err != nil {
		log.Println("Error read responce server: ", err)
	}
	log.Println(responce)

	var command string
	fmt.Scan(&command)
	fmt.Fprintln(conn, command)

	switch command {
	case "upload":
		commands.Upload(conn)
		HandlerRequest(conn)
	case "download":
		commands.Download(conn)
		HandlerRequest(conn)
	case "all_files":
		commands.AllFiles(conn)
		HandlerRequest(conn)
	default:
		responce := bufio.NewReader(conn)
		request, err := responce.ReadString('\n')
		if err != nil {
			log.Println("Error read responce server: ", err)
		}
		log.Println(request)
		command_default := "Неверная команда"
		fmt.Fprintln(conn, command_default)
		HandlerRequest(conn)
	}
}
