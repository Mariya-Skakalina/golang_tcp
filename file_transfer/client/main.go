package main

import (
	"file_transfer/client/handler"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	TYPE := os.Getenv("TYPE")

	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()

	for {
		handler.HandlerRequest(conn)
	}
}
