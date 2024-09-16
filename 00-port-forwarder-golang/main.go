package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var (
	LISTEN_PORT = flag.String("listenport", "5000", "Port to listen on")
	TARGET_IP   = flag.String("forwardip", "127.0.0.1", "IP where the network traffic is supposed to be forwarded")
	TARGET_PORT = flag.String("forwardport", "3389", "Port to forward to")
)

func main() {
	flag.Parse()

	listenPort := *LISTEN_PORT
	targetIP := *TARGET_IP
	targetPort := *TARGET_PORT

	// Start listening on the specified port
	listener, err := net.Listen("tcp", ":"+listenPort)
	if err != nil {
		log.Fatalf("Error starting TCP listener: %v", err)
	}
	defer listener.Close()
	log.Printf("Listening on port %s and forwarding to %s:%s\n", listenPort, targetIP, targetPort)

	for {
		// Accept an incoming connection
		clientConn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Forward the connection to the target address
		go forwardConnection(clientConn, targetIP+":"+targetPort)
	}
}

func forwardConnection(clientConn net.Conn, targetAddr string) {
	// Connect to the target address
	targetConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Printf("Error connecting to target address %s: %v", targetAddr, err)
		clientConn.Close()
		return
	}

	// Bi-directional copy of data between client and target
	go io.Copy(targetConn, clientConn)
	go io.Copy(clientConn, targetConn)

	log.Printf("Forwarded connection from %s to %s\n", clientConn.RemoteAddr(), targetAddr)
}
