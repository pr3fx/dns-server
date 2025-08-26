package dns

import (
	"fmt"
	"net"
)

type RecordType uint16
const (
	RecordType_A RecordType = 1
	RecordType_AAAA RecordType = 28
	RecordType_NS RecordType = 2
)

func encodeDomainName(domain_name string) []byte {
	var encoded_domain = make([]byte, 1, len(domain_name)+2)
	var i int
	j := uint8(0)
	k := 0
	for i=0; i<len(domain_name); i++ {
		if domain_name[i] != "."[0] {
			encoded_domain = append(encoded_domain, domain_name[i])
			j++
		} else {
			encoded_domain[k] = j
			encoded_domain = append(encoded_domain, 0)
			k = i + 1
			j = 0
		}
	}
	encoded_domain[k] = j
	encoded_domain = append(encoded_domain, 0)

	return encoded_domain
}

func getEncodedDomainLen(buf []byte) (uint16, error) {
	var byte_count uint16 = 1
	j := uint8(buf[0])

	for j != 0 {
		byte_count += uint16(j) + 1
		if len(buf) < int(byte_count) {
			return 0, fmt.Errorf(`Encoded domain malformed or bytestream truncated`)
		}
		j = buf[byte_count-1]
	}

	return byte_count, nil
}


// Public DNS server functions
func NewDNSMessageResponse(message_stream []byte) (DNSMessage, error) {
	received_msg, err := ParseMessage(message_stream)
	if err != nil {
		return DNSMessage{}, err
	}
	// Construct response header
	rsp_header := DNSHeader{
		id:received_msg.header.GetID(), // mimic received ID field
	}
	rsp_header.SetQR(1) // set response type to response
	rsp_header.SetOPCODE(received_msg.header.GetOPCODE())
	rsp_header.SetAA(0)
	rsp_header.SetTC(0)
	rsp_header.SetRD(received_msg.header.GetRD())
	rsp_header.SetRA(0)
	rsp_header.SetZ(0)
	if received_msg.header.GetOPCODE() == 0 {
		rsp_header.SetRCODE(0) // no error
	} else {
		rsp_header.SetRCODE(4) // not implemented
	}
	// Set to 0 for now (ie. nothing in question or answer sections)
	rsp_header.SetQDCOUNT(0)
	rsp_header.SetANCOUNT(0)
	rsp_header.SetNSCOUNT(0)
	rsp_header.SetARCOUNT(0)

	rsp_message := DNSMessage{header:rsp_header}

	// Add questions and answers
	rsp_message.AddQuestion(NewDNSQuestion("codecrafters.io", RecordType_A, 1))
	dns_ans_type, err := NewTypeA_Answer(net.IPv4(8,8,8,8))
	if err != nil {
		return DNSMessage{}, fmt.Errorf(`Encountered error while creating dns_ans_type: %v`, err)
	}
	rsp_message.AddAnswer(NewDNSAnswer("codecrafters.io", 1, 60, dns_ans_type))

	return rsp_message, nil
}
