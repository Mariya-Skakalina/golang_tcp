package main

import (
	"fmt"
	"io"
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

	// Выбираем файл для загрузки (картинка)
	filePath := "./image.jpeg" // Путь к картинке
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Читаем содержимое файла в буфер
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// Отправляем содержимое файла на сервер
	_, err = conn.Write(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Файл успешно загружен на сервер")
}

// Отправляет имя файла и его содержимое серверу
func uploadFile(conn net.Conn, fileName string) {
	// Отправьте имя файла серверу
	_, err := fmt.Fprintf(conn, "%s\n", fileName)
	if err != nil {
		log.Println("Failed to send file name:", err)
		return
	}

	// Откройте файл и отправьте его содержимое серверу
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(conn, file)
	if err != nil {
		log.Println("Failed to send file:", err)
		return
	}

	log.Println("File", fileName, "uploaded successfully.")
}
