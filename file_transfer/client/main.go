package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Устанавливаем соединение с сервером
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		fmt.Print("Введите команду (upload, download, exit, all_file): ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		fmt.Fprintf(conn, "%s\n", command)

		switch command {
		case "upload":
			upload(conn)
		case "download":
			fmt.Println("Выберите файл для скачивания")
			// Здесь должен быть код для скачивания файла с сервера
		case "all_files":
			all_files(conn)
		case "exit":
			fmt.Println("Завершение работы клиента")
			return
		default:
			fmt.Println("Неизвестная команда. Пожалуйста, выберите команду из списка.")
		}
	}
}

func upload(conn net.Conn) {
	defer conn.Close()
	fmt.Print("Введите имя для файла: ")
	name_file, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	name_file = strings.TrimSpace(name_file)
	fmt.Fprintf(conn, "%s\n", name_file)

	fmt.Print("Выберите файл для загрузки: ")
	filePath, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка получения информации о файле:", err)
		return
	}

	buffer := make([]byte, fileInfo.Size())
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println("Ошибка отправки файла на сервер:", err)
		return
	}

	// Получаем подтверждение от сервера о получении файла
	confirmation, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения подтверждения:", err)
		return
	}
	fmt.Println("Ответ сервера:", confirmation)
}

func all_files(conn net.Conn) {
	// Чтение количества файлов
	countStr, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Ошибка чтения количества файлов:", err)
		return
	}
	count, err := strconv.Atoi(strings.TrimSpace(countStr))
	if err != nil {
		log.Println("Ошибка преобразования количества файлов:", err)
		return
	}

	// Чтение имен файлов
	for i := 0; i < count; i++ {
		name, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Ошибка чтения имени файла:", err)
			return
		}
		fmt.Println("Имя файла:", strings.TrimSpace(name))
	}
}
