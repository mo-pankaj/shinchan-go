package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"strconv"
)

func doSomething(conn net.Conn) {
	// conn.Read is blocking
	n, err := conn.Read(make([]byte, 1024))
	if err != nil {
		log.Fatalf("error reading. error: %v" + err.Error())
	}
	slog.Info("number of bytes" + strconv.Itoa(n))
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World\r\n"))
	conn.Close()
}

func main() {
	fmt.Print("Server starting")

	// tcp works on a connection
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to get listener. error: %v" + err.Error())
	}

	slog.Info("Server listening", "listner", listener)

	// infinte loop to accept connections
	for {

		// tcp accept function is a blocking function
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("unable to get listener. error: %v" + err.Error())
		}

		slog.Info("connection waiting for accept. ", "conn", conn)

		// making async
		go doSomething(conn)
	}

	log.Fatalf("server clossed")
}
