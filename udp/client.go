package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// Define command-line flags
	var serverAddr string
	var message string
	var requestsCount int
	var showHelp bool
	var sleepTime int

	flag.StringVar(&serverAddr, "server", "", "UDP server address")
	flag.StringVar(&message, "message", "", "UDP message")
	flag.IntVar(&requestsCount, "count", 0, "Number of UDP requests to send")
	flag.IntVar(&sleepTime, "sleep", 100, "Sleep time in microseconds between UDP requests")
	flag.BoolVar(&showHelp, "h", false, "Show help")

	// Set custom usage information
	flag.Usage = func() {
		fmt.Printf("Usage: %s -server <server_address> -message <udp_message> -count <num_requests> -sleep <sleep_time_microseconds>\n", flag.CommandLine.Name())
		flag.PrintDefaults()
	}

	// Parse command line arguments
	flag.Parse()

	// Check if help flag is set
	if showHelp || flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	// Check if required flags are provided
	if serverAddr == "" || message == "" || requestsCount <= 0 {
		fmt.Println("Error: Missing required flags. Use -h or --help for usage information.")
		return
	}

	// Create a channel to receive results
	resultChannel := make(chan bool, requestsCount)

	// Create a wait group to wait for all requests to finish
	var wg sync.WaitGroup

	// Record start time
	startTime := time.Now()

	// Start the UDP requests
	for i := 0; i < requestsCount; i++ {
		wg.Add(1)
		go sendUDPRequest(serverAddr, message, resultChannel, &wg)
		time.Sleep(time.Duration(sleepTime) * time.Microsecond)
	}

	// Wait for all requests to finish
	wg.Wait()

	// Record end time
	endTime := time.Now()

	// Close the result channel
	close(resultChannel)

	// Count successful connections
	successfulConnections := 0
	for result := range resultChannel {
		if result {
			successfulConnections++
		}
	}

	// Print the number of successful connections and elapsed time
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Number of successful connections:", successfulConnections)
	fmt.Printf("Elapsed time: %v\n", elapsedTime)
}

func sendUDPRequest(serverAddr, message string, resultChannel chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		resultChannel <- false
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error dialing UDP:", err)
		resultChannel <- false
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error writing to UDP:", err)
		resultChannel <- false
		return
	}

	// Successfully connected
	resultChannel <- true
}
