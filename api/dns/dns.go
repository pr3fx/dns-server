package dns

type RecordType uint16
const (
	RecordType_A RecordType = 1
	RecordType_AAAA RecordType = 28
	RecordType_NS RecordType = 2
)
