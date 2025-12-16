package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	fmt.Println("Listener: ", listener)

	if err != nil {
		log.Fatal(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			buf := make([]byte, 1024)
			_, err := c.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			conn.Write([]byte("Hello from TCP server"))
			fmt.Print("Message Received: ", string(buf))
			c.Close()
		}(conn)
	}
}
