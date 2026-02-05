package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected")

	reader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)
	connWriter := bufio.NewWriter(conn)

	for {
		// Read from console and send to server
		fmt.Print("You: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		connWriter.WriteString(msg + "\n")
		connWriter.Flush()

		if msg == "Over" {
			break
		}

		// Read from server
		response, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}

		response = strings.TrimSpace(response)
		fmt.Println("Server says:", response)

		if response == "Over" {
			break
		}
	}
}
