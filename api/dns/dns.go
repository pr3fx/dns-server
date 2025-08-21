package dns

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
