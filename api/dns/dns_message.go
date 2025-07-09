package dns

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)


type DNSMessage struct {
	header DNSHeader
	question []DNSQuestion
}

func (message *DNSMessage) SetHeader(header DNSHeader) {
	(*message).header = header
}

func (message *DNSMessage) AddQuestion(question DNSQuestion) {
	(*message).question = append((*message).question, question)
	// Increment QDCOUNT header field
	(*message).header.SetQDCOUNT((*message).header.qdcount + 1)
}

func (message DNSMessage) Serialize() []byte {
	header_buf := message.header.Serialize()
	question_buf_cap := uint16(0)
	for _, q := range message.question {
		question_buf_cap += q.GetByteLen()
	}
	question_buf := make([]byte, 0, question_buf_cap)
	for _, q := range message.question {
		question_buf = append(question_buf, q.Serialize()...)
	}
	if uint16(len(question_buf)) != question_buf_cap {
		log.Warning(
			fmt.Sprintf("Expected DNSQuestion buffer length and actual length differ: " +
                "%v bytes expected, %v bytes actual",
				question_buf_cap,
				len(question_buf),
			))
	}

	buf := make([]byte, 0, len(header_buf) + len(question_buf))
	buf = append(buf, header_buf...)
	buf = append(buf, question_buf...)
	return buf
}
