package dns

import (
	"encoding/binary"
)


type DNSQuestion struct {
	qname []byte
	qtype RecordType
	qclass uint16
}

func (question *DNSQuestion) SetQNAME(domain_name string) {
	(*question).qname = make([]byte, 1, len(domain_name)+2)
	// Encode domain name string
	var i int
	j := uint8(0)
	k := 0
	for i=0; i<len(domain_name); i++ {
		if domain_name[i] != "."[0] {
			(*question).qname = append((*question).qname, domain_name[i])
			j++
		} else {
			(*question).qname[k] = j
			(*question).qname = append((*question).qname, 0)
			k = i + 1
			j = 0
		}
	}
	(*question).qname[k] = j
	(*question).qname = append((*question).qname, 0)
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
