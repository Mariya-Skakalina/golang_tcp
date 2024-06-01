package commands

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// const (
// 	chunkSize = 4096 // Размер каждой части файла
// )

func Upload(conn net.Conn) {

	info, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println("Error text to server")
	}
	fmt.Println("Answer to server: ", info)

	if _, err = conn.Write([]byte("Hello from client\n")); err != nil {
		log.Println("error send qnswer from server", err)
	}

	// Запрос имени файла и расширения у пользователя
	// fmt.Print("Введите имя для файла: ")
	// nameFile, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// nameFile = strings.TrimSpace(nameFile)
	// fmt.Fprintf(conn, "%s\n", nameFile)

	// fmt.Print("Введите расширение файла: ")
	// fileExtension, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// fileExtension = strings.TrimSpace(fileExtension)
	// fmt.Fprintf(conn, "%s\n", fileExtension)

	// Запрос пути к файлу у пользователя
	// fmt.Print("Выберите файл для загрузки: ")
	// filePath, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// filePath = strings.TrimSpace(filePath)

	// file, err := os.Open(filePath)
	// if err != nil {
	// 	fmt.Println("Ошибка открытия файла:", err)
	// 	return
	// }
	// defer file.Close()

	// // Создание буфера для чтения файла
	// buffer := make([]byte, chunkSize)
	// writer := bufio.NewWriterSize(conn, chunkSize)

	// for {
	// 	n, err := file.Read(buffer)
	// 	fmt.Println(n)
	// 	if err != nil && err != io.EOF {
	// 		log.Fatal(err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	if err = binary.Write(writer, binary.LittleEndian, uint32(n)); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if _, err = conn.Write(buffer[:n]); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// fmt.Println("Файл загружен, ожидаем подтверждения от сервера...")

	// confirmation, err := bufio.NewReader(conn).ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Ошибка чтения подтверждения:", err)
	// 	return
	// }
	// fmt.Println("Ответ сервера:", confirmation)
}
