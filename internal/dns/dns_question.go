package dns

import (
	"encoding/binary"
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
