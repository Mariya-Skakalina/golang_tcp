package commander

import (
	"encoding/json"
	"io"
	"log"
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

func NewFile(id uuid.UUID, name string, path string) File {
	return File{id, name, path}
}

func ReadFiles() (Files, error) {
	var files Files
	file, err := os.Open("../files.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Если файл не существует, создаем его и записываем начальное значение
			file, err = os.Create("../files.json")
			if err != nil {
				return files, err
			}
			defer file.Close()

			initialData := Files{Files: []File{}}
			byteValue, err := json.MarshalIndent(initialData, "", "	")
			if err != nil {
				return files, err
			}

			if _, err := file.Write(byteValue); err != nil {
				return files, err
			}

			return initialData, nil
		} else {
			return files, err
		}
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return files, err
	}

	if len(byteValue) == 0 {
		return Files{Files: []File{}}, nil
	}

	if err = json.Unmarshal(byteValue, &files); err != nil {
		return files, err
	}

	return files, nil
}

func WriteFile(id uuid.UUID, name string, path string) {
	existingFiles, err := ReadFiles()
	if err != nil {
		log.Fatal(err)
	}
	newFile := NewFile(id, name, path)
	existingFiles.Files = append(existingFiles.Files, newFile)

	jsonData, err := json.MarshalIndent(existingFiles, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	if err = os.WriteFile("../files.json", jsonData, 0666); err != nil {
		log.Fatal(err)
	}
}
