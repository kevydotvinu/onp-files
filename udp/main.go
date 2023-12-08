package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	_, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading from UDP:", err)
		return
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	response := fmt.Sprintf("Hostname: %s\n", hostname)
	_, err = conn.WriteToUDP([]byte(response), addr)
	if err != nil {
		fmt.Println("Error writing to UDP:", err)
		return
	}
}

func main() {
	port := flag.Int("port", 8080, "Port for the server to listen on")
	flag.Parse()

	// Create a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	// Listen for UDP connections
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening for UDP:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Server is running on UDP port %d\n", *port)

	for {
		handleUDPConnection(conn)
	}
}
