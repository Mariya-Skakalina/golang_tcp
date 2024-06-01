package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {

		info := bufio.NewReader(os.Stdin)
		if err != nil {
			log.Println("Ошибка чтения команд от сервера")
		}
		fmt.Println(info.Size())
		var command string
		fmt.Scan(&command)
		if _, err = conn.Write([]byte(command)); err != nil {
			log.Println("Ошибка передачи команды; ", err)
		}
		// command = strings.TrimSpace(command)
		// fmt.Fprintf(conn, "%s\n", command)

		switch command {
		case "upload":
			// commands.Upload(conn)
			// continue
			fmt.Println("upload")
		case "download":
			fmt.Println("Выберите файл для скачивания")
			continue
			// Здесь должен быть код для скачивания файла с сервера
		case "all_files":
			// all_files(conn)
			continue
		case "exit":
			fmt.Println("Завершение работы клиента")
			return
		default:
			fmt.Println("Неизвестная команда. Пожалуйста, выберите команду из списка.")
			continue
		}
	}
}

// func all_files(conn net.Conn) {

// 	// Чтение количества файлов
// 	// countStr, err := bufio.NewReader(conn).ReadString('\n')
// 	// if err != nil {
// 	// 	log.Println("Ошибка чтения количества файлов:", err)
// 	// 	return
// 	// }
// 	// count, err := strconv.Atoi(strings.TrimSpace(countStr))
// 	// if err != nil {
// 	// 	log.Println("Ошибка преобразования количества файлов:", err)
// 	// 	return
// 	// }

// 	// // Чтение имен файлов
// 	// for i := 0; i < count; i++ {
// 	// 	name, err := bufio.NewReader(conn).ReadString('\n')
// 	// 	if err != nil {
// 	// 		if err == io.EOF {
// 	// 			break
// 	// 		}
// 	// 		log.Println("Ошибка чтения имени файла:", err)
// 	// 		return
// 	// 	}
// 	// 	fmt.Println("Имя файла:", strings.TrimSpace(name))
// 	// }
// }
