package dns

import (
	"fmt"
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
