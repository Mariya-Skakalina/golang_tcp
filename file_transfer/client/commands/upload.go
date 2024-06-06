package commands

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func Upload(conn net.Conn) {
	request := bufio.NewReader(conn)
	responce, err := request.ReadString('\n')
	if err != nil {
		log.Println("Ошибка записи имени файла: ", err)
	}
	fmt.Println(responce)

	readerf := bufio.NewReader(os.Stdin)
	nameFile, err := readerf.ReadString('\n')
	if err != nil {
		log.Println("Ошибка чтения имени файла: ", err)
	}
	nameFile = strings.TrimSpace(nameFile)
	fmt.Fprintln(conn, nameFile)

	requestPath := bufio.NewReader(conn)
	responcePath, err := requestPath.ReadString('\n')
	if err != nil {
		log.Println("Ошибка сохранения файла или файл не найден: ", err)
	}
	fmt.Println(responcePath)

	filePath, err := bufio.NewReader(os.Stdin).ReadString('\n')
	filePath = strings.TrimSpace(filePath)
	if err != nil {
		log.Println("Ошибка пути файла: ", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Ошибка открытия файла: ", err)
		os.Exit(1)
	}
	size, _ := file.Stat()
	sizeFile := size.Size()
	defer file.Close()

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
}
