package commander

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func AllFiles(conn net.Conn) {
	var files Files
	file, err := os.Open("../files.json")
	if err != nil {
		log.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Println("Ошибка чтения файла:", err)
	}
	if err = json.Unmarshal(byteValue, &files); err != nil {
		log.Fatal(err)
	}
	byteFiles, err := json.Marshal(files)
	if err != nil {
		log.Println(err)
	}
	for _, v := range files.Files {
		fmt.Println(v.Name)
	}
	fmt.Println(string(byteFiles))

	// // Отправляем данные клиенту
	// if _, err := conn.Write(byteFiles); err != nil {
	// 	log.Println("Ошибка отправки данных:", err)
	// }
}
