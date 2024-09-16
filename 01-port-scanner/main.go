package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
)

var (
	TARGET_IP     = flag.String("ip", "127.0.0.1", "IP to scan")
	PORTS_TO_SCAN = flag.Int("scanuntil", 65535, "Port to scan until from 0")
)

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", *TARGET_IP, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	flag.Parse()
	// number of simultaneous workers
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < *PORTS_TO_SCAN; i++ {
			ports <- i
		}
	}()

	for i := 1; i < *PORTS_TO_SCAN; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	fmt.Printf("Open ports for %s\n", *TARGET_IP)
	for _, port := range openports {
		fmt.Printf("%d\n", port)
	}
}
