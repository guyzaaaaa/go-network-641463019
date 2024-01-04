package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("641463019 ธนภัทร อภิวงค์")

	fmt.Print("Connecting to server...\n")

	fmt.Print("กรอก username: ")
	username := getUserInput()

	fmt.Print("กรอก password: ")
	password := getUserInput()

	data := fmt.Sprintf("%s:%s", strings.TrimSpace(username), strings.TrimSpace(password))

	conn, err := net.Dial("tcp", "localhost:87")
	if handleError(err, "Error connecting to server") {
		return
	}
	defer conn.Close()

	if _, err := conn.Write([]byte(data)); handleError(err, "Error sending data to server") {
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if handleError(err, "Error receiving data from server") {
		return
	}

	fmt.Println("Server response:", string(buffer[:n]))
}

func getUserInput() string {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return input
}

func handleError(err error, message string) bool {
	if err != nil {
		fmt.Println(message, ":", err)
		return true
	}
	return false
}
