package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server started")
	fmt.Println("Waiting for a client ...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client accepted")

	reader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)
	connWriter := bufio.NewWriter(conn)

	for {
		// Read from client
		msg, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}

		msg = strings.TrimSpace(msg)
		fmt.Println("Client says:", msg)

		if msg == "Over" {
			break
		}

		// Read from console and send to client
		fmt.Print("You: ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)

		connWriter.WriteString(response + "\n")
		connWriter.Flush()

		if response == "Over" {
			break
		}
	}
}
