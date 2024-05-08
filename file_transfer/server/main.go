package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var (
	host     = "localhost"
	port     = "8000"
	typeServ = "tcp"
	mu       sync.Mutex
)

func init() {
	fmt.Println("Woo")
}

func main() {
	// Проверяем или создаем директорию для сохранения файлов, если ее нет
	_, err := os.Stat("uploads")
	if os.IsNotExist(err) {
		err := os.Mkdir("uploads", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

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

type File struct {
	ID   uuid.UUID
	Name string
	Path string
}

var files []File

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
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
			upload(conn, reader)
		case "download":
			download(conn, reader)
		case "all_files":
			all_files(conn)
		default:
			log.Println("Неизвестная команда:", command)
			return
		}
	}
}

func upload(conn net.Conn, reader *bufio.Reader) {
	var id = uuid.New()

	name_file, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Ошибка чтения имени файла:", err)
		return
	}
	name_file = strings.TrimSpace(name_file)

	if name_file == "" {
		log.Println("Имя файла не может быть пустым")
		return
	}

	filePath := fmt.Sprintf("uploads/%s.jpeg", id.String())
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	var file_instance = File{
		ID:   id,
		Name: name_file,
		Path: filePath,
	}

	mu.Lock()
	files = append(files, file_instance)
	mu.Unlock()

	// Копируем данные из соединения в файл
	_, err = io.Copy(file, reader)
	if err != nil {
		log.Println("Ошибка копирования данных:", err)
		return
	}
	conn.Close()

	confirmation := "Файл успешно загружен\n"
	if _, err := conn.Write([]byte(confirmation)); err != nil {
		log.Println("Ошибка отправки подтверждения:", err)
		return
	}
	log.Println("Подтверждение отправлено клиенту")

	log.Println("Файл получен и сохранен:", filePath)
}

func download(conn net.Conn, reader *bufio.Reader) {

}

func all_files(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	// Отправляем количество файлов
	count := len(files)
	conn.Write([]byte(fmt.Sprintf("%d\n", count)))

	// Отправляем имена файлов
	for _, file := range files {
		conn.Write([]byte(file.Name + "\n"))
	}
}
