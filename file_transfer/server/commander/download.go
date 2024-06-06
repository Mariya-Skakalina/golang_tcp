package commander

import (
	"fmt"
	"net"
)

func Download(conn net.Conn) {
	fmt.Println(conn)
	fmt.Println("Hello download")
}
