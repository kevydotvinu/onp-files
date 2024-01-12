package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
)

var requestCount int
var mu sync.Mutex

func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	_, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Println("Error reading from UDP:", err)
		return
	}

	mu.Lock()
	requestCount++
	totalRequests := requestCount
	mu.Unlock()

	fmt.Printf("Source IP: %s, Source Port: %d - Total Requests: %d\n", addr.IP, addr.Port, totalRequests)
}

func main() {
	var port int
	var showHelp bool

	flag.IntVar(&port, "port", 0, "Port for the server to listen on")
	flag.BoolVar(&showHelp, "h", false, "Show help")

	flag.Usage = func() {
		fmt.Printf("Usage: %s -port <port_number>\n", flag.CommandLine.Name())
		flag.PrintDefaults()
	}

	flag.Parse()

	if showHelp || port <= 0 {
		flag.Usage()
		return
	}

	// Create a UDP address
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println("Error resolving UDP address:", err)
		return
	}

	// Listen for UDP connections
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Println("Error listening for UDP:", err)
		return
	}
	defer conn.Close()

	log.Printf("Server is running on UDP port %d\n", port)

	// Block indefinitely to handle UDP connections
	for {
		handleUDPConnection(conn)
	}
}
