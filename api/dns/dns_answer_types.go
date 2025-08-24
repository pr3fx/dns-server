package dns

import (
	"net"
)

// ----------------- ANSWER INTERFACES -----------------
type DNSAnswerRecordType interface {
	getRecordType() RecordType
}

type DNSAnswerRDATA interface {
	serializeToBigEndianBytes() ([]byte, uint32)
}

// composite interface
type DNSAnswerIntf interface {
	DNSAnswerRecordType
	DNSAnswerRDATA
}
// -----------------------------------------------------

// TYPE A ANSWER
type typeA_Answer struct {
	ipv4_addr net.IP
}

func NewTypeA_Answer(ipv4_addr net.IP) typeA_Answer {
	return typeA_Answer{ ipv4_addr:ipv4_addr }
}

func (data typeA_Answer) serializeToBigEndianBytes() ([]byte, uint32) {
	return []byte(data.ipv4_addr), uint32(len(data.ipv4_addr))
}

func (answer typeA_Answer) getRecordType() RecordType {
	return RecordType_A
}
