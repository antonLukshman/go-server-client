package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type a message:")

	// Read input from the user
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	// Send message to the server
	fmt.Fprintf(conn, message)

	// Read response from the server
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Server response:", response)
}
