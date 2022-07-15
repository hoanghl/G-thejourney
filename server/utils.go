package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

// type envVar

func getEnv() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("STACK")
	fmt.Println(val)

	val = os.Getenv("DATABASE")
	fmt.Println(val)

}

func handleIncomingConn(conn net.Conn) {
	buff := make([]byte, 1024)
	conn.Read(buff)

	fmt.Printf("Input from client\n: %s", string(buff[:]))

	conn.Write([]byte("Done here, bro"))
	conn.Close()
}
