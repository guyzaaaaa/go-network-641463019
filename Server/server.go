package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// กำหนด username และ passwordเป็นตามนี้
	validUsername := "std1"
	validPassword := "p@ssw0rd"

	// รับข้อมูลจาก Client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	clientData := string(buffer[:n])

	if clientData == fmt.Sprintf("%s:%s", validUsername, validPassword) {

		conn.Write([]byte("Hello\n"))
	} else {

		conn.Write([]byte("Invalid credentials\n"))
	}
}

func main() {
	fmt.Println("Server is starting...")

	// เปิด port 12345 เพื่อรองรับการเชื่อมต่อ
	ln, err := net.Listen("tcp", ":87")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	for {
		// รอรับการเชื่อมต่อจาก Client
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// เริ่มต้นการจัดการเชื่อมต่อ
		go handleConnection(conn)
	}
}
