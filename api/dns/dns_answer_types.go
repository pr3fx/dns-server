package dns

import (
	"net"
	"fmt"
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

func NewTypeA_Answer(ipv4_addr net.IP) (typeA_Answer, error) {
	// ensure ipv4_addr is 4 bytes, not 16
	ipv4_addr_4 := ipv4_addr.To4()
	if ipv4_addr_4 == nil {
		return typeA_Answer{}, fmt.Errorf(`IPv4 address is not valid: %v`, ipv4_addr)
	}
	return typeA_Answer{ ipv4_addr:ipv4_addr_4 }, nil
}

func (data typeA_Answer) serializeToBigEndianBytes() ([]byte, uint32) {
	return []byte(data.ipv4_addr), uint32(len(data.ipv4_addr))
}

func (answer typeA_Answer) getRecordType() RecordType {
	return RecordType_A
}
