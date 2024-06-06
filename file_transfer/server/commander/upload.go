package commander

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
)

func Upload(conn net.Conn) {
	_, err := os.Stat("uploads")
	if os.IsNotExist(err) {
		err := os.Mkdir("uploads", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	id := uuid.New()

	infoUpload := "Введите имя файла для сохранения: \n"
	if _, err := conn.Write([]byte(infoUpload)); err != nil {
		log.Println("Ошибка записи имени файла: ", err)
	}

	name := bufio.NewReader(conn)
	nameFile, err := name.ReadString('\n')
	if err != nil {
		log.Println("Ошибка при сохранении имени файла")
	}
	nameFile = strings.TrimSpace(nameFile)

	uploadFile := "Введите путь до файла \n"
	if _, err = conn.Write([]byte(uploadFile)); err != nil {
		log.Println("Такого файла не существует или другая ошибка ", err)
	}

	filePath := fmt.Sprintf("uploads/%s.jpeg", id.String())
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Ошибка при создании файла: ", err)
	}
	defer file.Close()

	WriteFile(id, nameFile, filePath)
	// fmt.Println(newFile)

	done := make(chan error)

	writer1 := bufio.NewWriter(file)
	reader1 := bufio.NewReader(conn)
	go func() {
		_, err := io.CopyN(writer1, reader1, 121978)
		done <- err
	}()
	err = <-done
	if err != nil && err != io.EOF {
		fmt.Println("Error during copy:", err)
	} else {
		fmt.Println("Copy completed successfully")
	}
}
