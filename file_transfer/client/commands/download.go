package commands

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

func Download(conn net.Conn) {
	reader, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println(err)
	}

	id := uuid.New()

	fmt.Print(reader)
	readerName := bufio.NewReader(os.Stdin)
	nameFile, err := readerName.ReadString('\n')
	if err != nil {
		log.Println("Ошибка чтения имени файла: ", err)
	}
	nameFile = strings.TrimSpace(nameFile)
	fmt.Fprintln(conn, nameFile)
	// readerResponce := bufio.NewReader(os.Stdin)
	// responce, err := readerResponce.ReadString('\n')
	// if err != nil {
	// 	log.Println("Ошибка чтения имени файла: ", err)
	// }

	_, err = os.Stat("uploads")
	if os.IsNotExist(err) {
		err := os.Mkdir("uploads", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	filePath := fmt.Sprintf("uploads/%s.jpeg", id.String())
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Ошибка при создании файла: ", err)
	}
	defer file.Close()
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
	}

}
