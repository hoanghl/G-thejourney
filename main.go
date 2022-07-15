package main

import (
	"flag"
	"fmt"
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

const PORT = 64321
const portSecond = 22

var (
	target string
	port   int
)

func init() {
	flag.StringVar(&target, "target", "", "target (<host>:<port>)")
	flag.IntVar(&port, "port", portSecond, "port")
}

func main() {

	incomming, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatal("Cannot start server at %d: %v", PORT, err)
		return
	}

	// accept connection
	client, err := incomming.Accept()
	if err != nil {
		log.Fatal("Conenction from client error: %v", err)
		return
	}

	fmt.Println(target)

	target, err := net.Dial("tcp", fmt.Sprintf(":%d", portSecond))
	if err != nil {
		log.Fatal("Dial got error: %v", err)
		return
	}
	go func() {
		io.Copy(target, client)
	}()
	go func() {
		io.Copy(client, target)
	}()
}
