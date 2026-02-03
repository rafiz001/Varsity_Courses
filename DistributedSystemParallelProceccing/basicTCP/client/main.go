package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Trying...")
		return
	}

	fmt.Println("Socket connected")

	conn.Close()
}
