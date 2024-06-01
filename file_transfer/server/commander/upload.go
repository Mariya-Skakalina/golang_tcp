package commander

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// const (
// 	chunkSize = 4096 // Размер каждой части файла
// )

// var mu sync.Mutex

// type File struct {
// 	ID   uuid.UUID
// 	Name string
// 	Path string
// }

// var files []File

func Upload(conn net.Conn, reader *bufio.Reader) {

	// Проверяем или создаем директорию для сохранения файлов, если ее нет
	_, err := os.Stat("uploads")
	if os.IsNotExist(err) {
		err := os.Mkdir("uploads", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	info := "Ответ от сервера\n"
	if _, err = conn.Write([]byte(info)); err != nil {
		log.Println("Ошибка на стороне сервера: ", err)
		return
	}

	client_info, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Ошибка чтания от клиента")
	}
	fmt.Println(client_info)

	// Генерация UUID для файла
	// id := uuid.New()

	// // Запрашиваем имя для файла
	// nameFile, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Println("Ошибка чтения имени файла:", err)
	// 	return
	// }
	// nameFile = strings.TrimSpace(nameFile)
	// if nameFile == "" {
	// 	log.Println("Имя файла не может быть пустым")
	// 	return
	// }

	// // Запрашиваем расширение для файла
	// fileExtension, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Println("Ошибка чтения расширения файла:", err)
	// 	return
	// }
	// fileExtension = strings.TrimSpace(fileExtension)
	// if fileExtension == "" {
	// 	log.Println("Расширение файла не может быть пустым")
	// 	return
	// }

	// // Создание файла
	// filePath := fmt.Sprintf("uploads/%s.jpeg", id.String())
	// file, err := os.Create(filePath)
	// if err != nil {
	// 	log.Println("Ошибка создания файла:", err)
	// 	return
	// }
	// defer file.Close()

	// fileInstance := File{
	// 	ID:   id,
	// 	Name: id.String(),
	// 	Path: filePath,
	// }

	// mu.Lock()
	// files = append(files, fileInstance)
	// mu.Unlock()

	// fmt.Println(reader.Size())

	// buffer := make([]byte, chunkSize)
	// var size uint32

	// if size > chunkSize {
	// 	log.Println("Ошибка: размер части файла превышает размер буфера")
	// 	// os.Remove(filePath) // Удаление файла при ошибке
	// 	fmt.Println("Ошибка: размер части файла превышает размер буфера")
	// 	return
	// }

	// if _, err := io.ReadFull(reader, buffer[:size]); err != nil {
	// 	log.Println("Ошибка чтения части файла: ", err)
	// 	// os.Remove(filePath) // Удаление файла при ошибке
	// 	fmt.Println("Ошибка чтения части файла")
	// 	return
	// }

	// if _, err := file.Write(buffer[:size]); err != nil {
	// 	log.Println("Ошибка записи части файла: ", err)
	// 	// os.Remove(filePath) // Удаление файла при ошибке
	// 	return
	// }
	// confirmation := "Файл успешно загружен\n"
	// if _, err := conn.Write([]byte(confirmation)); err != nil {
	// 	log.Println("Ошибка отправки подтверждения:", err)
	// 	return
	// }
	log.Println("Подтверждение отправлено клиенту")
}
