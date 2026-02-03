package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Server started on port 3000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			listener.Close()
			continue
		}
		fmt.Println("Connection is accepted from", conn.RemoteAddr())
		conn.Close()
	}

}
