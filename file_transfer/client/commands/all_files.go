package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/google/uuid"
)

type Files struct {
	Files []File `json:"files"`
}

type File struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Path string    `json:"path"`
}

func AllFiles(conn net.Conn) {
	var files Files
	file, err := os.Open("../files.json")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	if len(byteValue) == 0 {
		log.Println(err)
	}

	if err = json.Unmarshal(byteValue, &files); err != nil {
		log.Println(err)
	}

	for _, v := range files.Files {
		fmt.Println(v.Name)
	}

}
