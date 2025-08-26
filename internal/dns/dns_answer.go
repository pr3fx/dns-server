package dns

import (
	"encoding/binary"
	"fmt"
)


type DNSAnswer struct {
	name []byte
	ans_type RecordType
	class uint16
	ttl uint32
	rdlength uint32
	rdata []byte
}

func NewDNSAnswer(domain_name string, class uint16, ttl uint32, answer DNSAnswerIntf) DNSAnswer {
	record_type := answer.getRecordType()
	rdata, rdlength := answer.serializeToBigEndianBytes()

	new_dns_answer := DNSAnswer{}
	new_dns_answer.SetNAME(domain_name)
	new_dns_answer.SetTYPE(record_type)
	new_dns_answer.SetCLASS(class)
	new_dns_answer.SetTTL(ttl)
	new_dns_answer.SetRDLENGTH(rdlength)
	new_dns_answer.SetRDATA(rdata)
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

func (answer *DNSAnswer) GetByteLen() uint32 {
	return uint32(len((*answer).name) + len((*answer).rdata) + 12)
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
		buf[len(answer.name)+12+idx] = v
	}

	return buf
}


// Parsing functions
func ParseAnswers(answer_stream []byte, ANCOUNT uint16) ([]DNSAnswer, int, error) {
	stream_pos := 0
	var answer_list []DNSAnswer
	for i := 0; i < int(ANCOUNT); i++ {
		NAME_len, err := getEncodedDomainLen(answer_stream[stream_pos:])
		if err != nil {
			return nil, 0, err
		}
		NAME_end_pos := stream_pos + int(NAME_len)
		// Check if stream is big enough for TYPE, CLASS, TTL, RDLENGTH fields
		if len(answer_stream) < NAME_end_pos + 12 {
			return nil, 0, fmt.Errorf(`Bytestream is truncated`)
		}
		// Parse TYPE, CLASS, TTL, RDLENGTH
		TYPE_parsed := RecordType(binary.BigEndian.Uint16(answer_stream[NAME_end_pos:NAME_end_pos+2]))
		CLASS_parsed := binary.BigEndian.Uint16(answer_stream[NAME_end_pos+2:NAME_end_pos+4])
		TTL_parsed := binary.BigEndian.Uint32(answer_stream[NAME_end_pos+4:NAME_end_pos+8])
		RDLENGTH_parsed := binary.BigEndian.Uint32(answer_stream[NAME_end_pos+8:NAME_end_pos+12])
		// Check if RDATA can fit in the remaining stream length
		if len(answer_stream) < NAME_end_pos + 12 + int(RDLENGTH_parsed) {
			return nil, 0, fmt.Errorf(`Bytestream is truncated, cannot read RDATA`)
		}

		// Create DNSAnswer from parsed values
		parsed_answer := DNSAnswer{
			name:answer_stream[stream_pos:NAME_end_pos],
			ans_type:TYPE_parsed,
			class:CLASS_parsed,
			ttl:TTL_parsed,
			rdlength:RDLENGTH_parsed,
			rdata:answer_stream[NAME_end_pos+12:NAME_end_pos+12+int(RDLENGTH_parsed)],
		}
		answer_list = append(answer_list, parsed_answer)

		// Update stream_pos to start at the next block
		stream_pos = NAME_end_pos + 12 + int(RDLENGTH_parsed)
	}
	return answer_list, stream_pos, nil
}
