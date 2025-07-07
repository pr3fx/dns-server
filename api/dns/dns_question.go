package dns


type DNSQuestion struct {
	name []byte
	record_type uint16
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

func (question *DNSQuestion) SetRecordType(record_type uint16) {
	(*question).record_type = record_type
}

func (question *DNSQuestion) SetClass(class uint16) {
	(*question).class = class
}
