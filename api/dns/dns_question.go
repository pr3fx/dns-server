package dns

import (
	"encoding/binary"
)


type DNSQuestion struct {
	name []byte
	record_type RecordType
	class uint16
}

func (question *DNSQuestion) SetName(domain_name string) {
	(*question).name = make([]byte, 1, len(domain_name)+2)
	// Encode domain name string
	var i int
	j := uint8(0)
	k := 0
	for i=0; i<len(domain_name); i++ {
		if domain_name[i] != "."[0] {
			(*question).name = append((*question).name, domain_name[i])
			j++
		} else {
			(*question).name[k] = j
			(*question).name = append((*question).name, 0)
			k = i + 1
			j = 0
		}
	}
	(*question).name[k] = j
	(*question).name = append((*question).name, 0)
}

func (question *DNSQuestion) SetRecordType(record_type RecordType) {
	(*question).record_type = record_type
}

func (question *DNSQuestion) SetClass(class uint16) {
	(*question).class = class
}

func (question *DNSQuestion) GetByteLen() uint16 {
	return uint16(len((*question).name) + 4)
}

// Serialize to bytes (big-endian)
func (question DNSQuestion) Serialize() []byte {
	buf := make([]byte, len(question.name) + 4)
	// Serialize encoded domain
	for idx, v := range question.name {
		buf[idx] = v
	}
	// Serialize type and class
	binary.BigEndian.PutUint16(buf[len(question.name):], uint16(question.record_type))
	binary.BigEndian.PutUint16(buf[len(question.name)+2:], question.class)

	return buf
}
