// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// )

// var (
// 	host      = "localhost"
// 	port      = "8000"
// 	type_serv = "tcp"
// )

// func main() {
// 	listen, err := net.Listen(type_serv, host+":"+port)
// 	if err != nil {
// 		log.Fatal(err)
// 		os.Exit(1)
// 	}
// 	defer listen.Close()

// 	// Каналы для входящих и прерванных соединений, сообщений
// 	aconns := make(map[net.Conn]int)
// 	conns := make(chan net.Conn)
// 	dconns := make(chan net.Conn)
// 	msgs := make(chan string)
// 	i := 1

// 	go func() {
// 		for {
// 			conn, err := listen.Accept()
// 			if err != nil {
// 				log.Fatal(err)
// 				os.Exit(1)
// 			}
// 			conns <- conn
// 		}
// 	}()
// 	for {
// 		select {
// 		// Читаем входящие соединения
// 		case conn := <-conns:
// 			aconns[conn] = i
// 			i++
// 			// Как только у нас есть соединения мы начинаем читать из него
// 			go func(conn net.Conn, i int) {
// 				rd := bufio.NewReader(conn)
// 				for {
// 					m, err := rd.ReadString('\n')
// 					fmt.Println(m)
// 					if err != nil {
// 						break
// 					}
// 					msgs <- fmt.Sprintf("Client %v: %v", i, m)
// 				}
// 				// Конец чтения
// 				dconns <- conn
// 			}(conn, i)
// 		case msg := <-msgs:
// 			// Трансляция по всем подключениям
// 			for conn := range aconns {
// 				conn.Write([]byte(msg))
// 				fmt.Println([]byte(msg))
// 			}
// 		case dconn := <-dconns:
// 			log.Printf("Client %v is done", aconns[dconn])
// 			delete(aconns, dconn)

// 		}
// 	}

// }

// // var clientMessage string

// // func handlerRequest(conn net.Conn) {
// // 	defer conn.Close()

// // 	log.Printf("Client connected from %s\n", conn.RemoteAddr())

// // 	buffer := make([]byte, 1024)
// // 	n, err := conn.Read(buffer)
// // 	if err != nil {
// // 		log.Println("Read error:", err)
// // 		return
// // 	}
// // 	fmt.Println("hi")
// // 	clientMessage := string(buffer[:n])
// // 	responseStr := fmt.Sprintf("Your message is: %v.", clientMessage)
// // 	_, err = conn.Write([]byte(responseStr))
// // 	if err != nil {
// // 		log.Println("Write error:", err)
// // 	}
// // }
