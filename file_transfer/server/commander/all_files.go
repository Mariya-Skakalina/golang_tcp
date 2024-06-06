package commander

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func AllFiles(conn net.Conn) {
	files, err := ReadFiles()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files.Files {
		fmt.Println(file.Name)
		buffer := []byte(file.Name + "\n")

		// Проверка размера буфера перед отправкой
		bufferSize := len(buffer)

		if err := binary.Write(conn, binary.LittleEndian, uint32(bufferSize)); err != nil {
			log.Println("Ошибка записи размера буфера: ", err)
			return // Если произошла ошибка, выходим из функции
		}

		if _, err := conn.Write(buffer); err != nil {
			log.Println("Ошибка записи имени файла: ", err)
			return // Если произошла ошибка, выходим из функции
		}
	}
}
