package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	response := fmt.Sprintf("Hostname: %s\n", hostname)
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing to TCP:", err)
		return
	}
}

func main() {
	port := flag.Int("port", 8080, "Port for the server to listen on")
	flag.Parse()

	// Create a TCP address
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("Error resolving TCP address:", err)
		return
	}

	// Listen for TCP connections
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error listening for TCP:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server is running on TCP port %d\n", *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting TCP connection:", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}
