package dns

import (
	"encoding/binary"
	"fmt"
)


type DNSQuestion struct {
	qname []byte
	qtype RecordType
	qclass uint16
}

func NewDNSQuestion(domain_name string, record_type RecordType, class uint16) DNSQuestion {
	new_dns_question := DNSQuestion{}
	new_dns_question.SetQNAME(domain_name)
	new_dns_question.SetQTYPE(record_type)
	new_dns_question.SetQCLASS(class)
	return new_dns_question
}

func (question *DNSQuestion) SetQNAME(domain_name string) {
	(*question).qname = encodeDomainName(domain_name)
}

func (question *DNSQuestion) SetQTYPE(record_type RecordType) {
	(*question).qtype = record_type
}

func (question *DNSQuestion) SetQCLASS(class uint16) {
	(*question).qclass = class
}

func (question *DNSQuestion) GetByteLen() uint16 {
	return uint16(len((*question).qname) + 4)
}

// Serialize to bytes (big-endian)
func (question DNSQuestion) Serialize() []byte {
	buf := make([]byte, len(question.qname) + 4)
	// Serialize encoded domain
	for idx, v := range question.qname {
		buf[idx] = v
	}
	// Serialize type and class
	binary.BigEndian.PutUint16(buf[len(question.qname):], uint16(question.qtype))
	binary.BigEndian.PutUint16(buf[len(question.qname)+2:], question.qclass)

	return buf
}


// Parsing functions
func ParseQuestions(question_stream []byte, QDCOUNT uint16) ([]DNSQuestion, error) {
	if len(question_stream) < 5 {
		return nil, fmt.Errorf(`bytestream is %v bytes, expected 5 or more.`, len(question_stream))
	}

	stream_pos := 0
	var question_list []DNSQuestion
	for i := 0; i < int(QDCOUNT); i++ {
		QNAME_len, err := getQNAMElen(question_stream[stream_pos:])
		if err != nil {
			return nil, err
		}
		QNAME_end_pos := stream_pos + int(QNAME_len)
		// Check if stream is big enough for the QTYPE and QCLASS fields
		if len(question_stream) <  + QNAME_end_pos + 4 {
			return nil, fmt.Errorf(`Bytestream is truncated`)
		}

		QTYPE_parsed := RecordType(binary.BigEndian.Uint16(question_stream[QNAME_end_pos:QNAME_end_pos+2]))
		QCLASS_parsed := binary.BigEndian.Uint16(question_stream[QNAME_end_pos+2:QNAME_end_pos+4])

		// Create DNSQuestion struct from parsed values, append to slice
		parsed_question := DNSQuestion{
			qname:question_stream[stream_pos:QNAME_end_pos],
			qtype:QTYPE_parsed,
			qclass:QCLASS_parsed,
		}
		question_list = append(question_list, parsed_question)

		// Update stream position to start at the next block
		stream_pos = QNAME_end_pos + 4
	}

	return question_list, nil
}

func getQNAMElen(buf []byte) (uint16, error) {
	var byte_count uint16 = 1
	j := uint8(buf[0])

	for j != 0 {
		byte_count += uint16(j) + 1
		if len(buf) < int(byte_count) {
			return 0, fmt.Errorf(`QNAME malformed or bytestream truncated`)
		}
		j = buf[byte_count-1]
	}

	return byte_count, nil
}
