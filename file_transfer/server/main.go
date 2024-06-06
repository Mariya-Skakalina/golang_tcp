package main

import (
	"file_transfer/server/handler"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Ошибка загрузки .env: ", err)
		os.Exit(1)
	}

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	TYPE := os.Getenv("TYPE")

	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal("Проблемы при создании подключения: ", err)
		os.Exit(1)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("Проблемы при запуске сервера: ", err)
			os.Exit(1)
		}

		go handler.HandlerRequest(conn)
	}
}
