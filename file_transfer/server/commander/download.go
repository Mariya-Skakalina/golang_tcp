package commander

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func Download(conn net.Conn) {
	text := []byte("Выберите файл для загрузки: \n")
	if _, err := conn.Write(text); err != nil {
		log.Println(err)
	}
	reader, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println(err)
	}
	fileName := strings.TrimSpace(reader)
	file, err := ReadFiles()
	if err != nil {
		log.Println(err)
	}
	for _, v := range file.Files {
		if v.Name == fileName {
			file, err := os.Open(v.Path)
			if err != nil {
				log.Println(err)
			}
			defer file.Close()
			size, _ := file.Stat()
			sizeFile := size.Size()

			reader := bufio.NewReader(file)
			buffer := make([]byte, sizeFile)

			for {
				n, err := reader.Read(buffer)
				fmt.Println(n)
				if err != nil {
					if err == io.EOF {
						break // Конец файла
					}
					log.Println("Ошибка чтения куска файла: ", err)
					os.Exit(1)
				}
				_, err = conn.Write(buffer[:n])
				if err != nil {
					log.Println("Ошибка записи на сервер: ", err)
				}
			}
		} else {
			if _, err := conn.Write([]byte("Такого файла нет, посмотрите список доступных файлов с помощью команды \"all_files\"\n")); err != nil {
				log.Println(err)
			}
			return
		}
	}
}
