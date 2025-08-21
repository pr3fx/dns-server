package dns

import (
	"encoding/binary"
)


type DNSAnswer struct {
	name []byte
	ans_type RecordType
	class uint16
	ttl uint32
	rdlength uint32
	rdata []byte
}

func NewDNSAnswer(domain_name string) DNSAnswer {
	new_dns_answer := DNSAnswer{}
	return new_dns_answer
}

func (answer *DNSAnswer) SetNAME(domain_name string) {
	(*answer).name = encodeDomainName(domain_name)
}

func (answer *DNSAnswer) SetTYPE(record_type RecordType) {
	(*answer).ans_type = record_type
}

func (answer *DNSAnswer) SetCLASS(class uint16) {
	(*answer).class = class
}

func (answer *DNSAnswer) SetTTL(ttl uint32) {
	(*answer).ttl = ttl
}

func (answer *DNSAnswer) SetRDLENGTH(rdlength uint32) {
	(*answer).rdlength = rdlength
}

func (answer *DNSAnswer) SetRDATA(rdata []byte) {
	(*answer).rdata = rdata
}

func (answer DNSAnswer) Serialize() []byte {
	buf := make([]byte, len(answer.name)+len(answer.rdata)+12)
	// Serialize encoded domain name
	for idx, v := range answer.name {
		buf[idx] = v
	}
	// Serialize ans_type, class, ttl, rdlength
	binary.BigEndian.PutUint16(buf[len(answer.name):], uint16(answer.ans_type))
	binary.BigEndian.PutUint16(buf[len(answer.name)+2:], answer.class)
	binary.BigEndian.PutUint32(buf[len(answer.name)+4:], answer.ttl)
	binary.BigEndian.PutUint32(buf[len(answer.name)+8:], answer.rdlength)
	// Serialize encoded domain name
	for idx, v := range answer.rdata {
		buf[12+idx] = v
	}

	return buf
}
