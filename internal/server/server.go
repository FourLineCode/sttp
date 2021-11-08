package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func Run(addr, port string) {
	listener, err := net.Listen(addr, port)
	if err != nil {
		log.Panicln("Error while starting server", err.Error())
	}

	defer listener.Close()
	log.Println("\n\nServer running on port 8080")
	log.Println("Listening for connections")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error establishing connection", err.Error())
			conn.Close()
			continue
		}

		log.Printf("Connection extablished: %v\n", conn.RemoteAddr().String())

		go handle(conn)
	}
}

const MaxBufferSize = 1024

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		bytes := make([]byte, MaxBufferSize)
		n, err := conn.Read(bytes)

		if n == 0 || err == io.EOF {
			conn.Close()
			return
		}

		if err != nil {
			log.Println("Error while reading from connection", err.Error())
			conn.Close()
			return
		}

		fmt.Printf("Message from %v\n", conn.RemoteAddr().String())
		fmt.Println(strings.TrimSpace(string(bytes)))
	}
}
