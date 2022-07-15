package main

import (
	"fmt"
	"net"
)

func main() {
	const PORT = 65530
	listener, _ := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	fmt.Printf("Server is working at: localhost:%d", PORT)
	for {
		conn, _ := listener.Accept()

		go handleIncomingConn(conn)
	}
}
