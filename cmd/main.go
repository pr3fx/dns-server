package main

import (
	"net"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Attempting to start UDP server...")

	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		log.Error("Failed to resolve UDP address:", err)
		return
	}
	log.Info("UDP address successfully resolved")
	
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Error("Failed to bind to address:", err)
		return
	}
	log.Info("Bind to address successful")
	defer udpConn.Close()
	
	buf := make([]byte, 512)
	
	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			log.Error("Error receiving data:", err)
			break
		}
	
		receivedData := string(buf[:size])
		log.Error("Received %d bytes from %s: %s", size, source, receivedData)
	
		// Create an empty response
		response := []byte{}
	
		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			log.Error("Failed to send response:", err)
		}
	}
}
