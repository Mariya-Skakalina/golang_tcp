package handler

import (
	"bufio"
	"file_transfer/server/commander"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func HandlerRequest(conn net.Conn) {
	// Отправляет клиенту информацию
	info := "Напишите команду(upload, download, all_files): \n"
	if _, err := conn.Write([]byte(info)); err != nil {
		log.Println("errors transfer info: ", err)
		os.Exit(1)
	}

	// Получать информацию от клиента
	buffer := bufio.NewReader(conn)
	commandStart, err := buffer.ReadString('\n')
	if err != nil {
		// Обработка ошибки
		log.Println("Клиент отключился")
		conn.Close()
		return
	}
	commandStart = strings.TrimSpace(commandStart)
	switch commandStart {
	case "upload":
		commander.Upload(conn)
		HandlerRequest(conn)
	case "download":
		commander.Download(conn)
		HandlerRequest(conn)
	case "all_files":
		commander.AllFiles(conn)
		HandlerRequest(conn)
	default:
		info := "Неизвестная команда, повторите ввод: \n"
		if _, err := conn.Write([]byte(info)); err != nil {
			log.Println("errors transfer info: ", err)
			os.Exit(1)
		}
		buffer := bufio.NewReader(conn)
		commandDefault, err := buffer.ReadString('\n')
		if err != nil {
			// Обработка ошибки
			log.Println("Клиент отключился")
			conn.Close()
			return
		}
		fmt.Println(commandDefault)
	}
	// log.Println("Получено от клиента:", string(command_start[:len(command_start)-1]))
}
