package main

import (
	"net"
	log "github.com/sirupsen/logrus"
	"github.com/pr3fx/dns-server-go/api/dns"
	"fmt"
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
		log.Info(fmt.Sprintf("Received %d bytes from %s: %s", size, source, receivedData))

		// Create a DNS response (header and question)
		dns_header := dns.DNSHeader{}
		dns_header.SetID(1234)
		dns_header.SetQR(1) // Set response type

		dns_msg := dns.DNSMessage{}
		dns_msg.SetHeader(dns_header)
		dns_msg.AddQuestion(dns.NewDNSQuestion("codecrafters.io", dns.RecordType_A, 1))
		dns_answer_type, err := dns.NewTypeA_Answer(net.IPv4(8,8,8,8))
		if err != nil {
			fmt.Errorf(`Encountered error while creating dns_answer_type: %v`, err)
		}
		dns_msg.AddAnswer(dns.NewDNSAnswer("codecrafters.io", uint16(1), uint32(60), dns_answer_type))
		response := dns_msg.Serialize()

		_, err = udpConn.WriteToUDP(response, source)
		if err != nil {
			log.Error("Failed to send response:", err)
		}
	}
}
