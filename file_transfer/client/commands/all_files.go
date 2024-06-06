package commands

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func AllFiles(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		// Чтение размера файла
		var sizeFile uint32
		if err := binary.Read(reader, binary.LittleEndian, &sizeFile); err != nil {
			if err == io.EOF {
				break // Выходим из цикла, если соединение закрыто
			}
			log.Println("Ошибка чтения размера файла:", err)
			return
		}

		// Чтение имени файла до символа новой строки
		fileName := make([]byte, sizeFile)
		if _, err := io.ReadFull(reader, fileName); err != nil {
			log.Println("Ошибка чтения имени файла:", err)
			return
		}

		// Вывод имени файла
		fmt.Println(string(fileName))
	}
}
