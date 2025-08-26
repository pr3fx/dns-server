package dns

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)


type DNSMessage struct {
	header DNSHeader
	question []DNSQuestion
	answer []DNSAnswer
}

func (message *DNSMessage) SetHeader(header DNSHeader) {
	(*message).header = header
}

func (message *DNSMessage) AddQuestion(question DNSQuestion) {
	(*message).question = append((*message).question, question)
	// Increment QDCOUNT header field
	(*message).header.SetQDCOUNT((*message).header.qdcount + 1)
}

func (message *DNSMessage) AddAnswer(answer DNSAnswer) {
	(*message).answer = append((*message).answer, answer)
	// Increment ANCOUNT header field
	(*message).header.SetANCOUNT((*message).header.ancount + 1)
}

func (message DNSMessage) serializeQuestions() []byte {
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

	return question_buf
}

func (message DNSMessage) serializeAnswers() []byte {
	answer_buf_cap := uint32(0)
	for _, a := range message.answer {
		answer_buf_cap += a.GetByteLen()
	}
	answer_buf := make([]byte, 0, answer_buf_cap)
	for _, a := range message.answer {
		answer_buf = append(answer_buf, a.Serialize()...)
	}
	if uint32(len(answer_buf)) != answer_buf_cap {
		log.Warning(
			fmt.Sprintf("Expected DNSAnswer buffer length and actual length differ: " +
                "%v bytes expected, %v bytes actual",
				answer_buf_cap,
				len(answer_buf),
			))
	}

	return answer_buf
}

func (message DNSMessage) Serialize() []byte {
	header_buf := message.header.Serialize()
	question_buf := message.serializeQuestions()
	answer_buf := message.serializeAnswers()

	buf := make([]byte, 0, len(header_buf) + len(question_buf) + len(answer_buf))
	buf = append(buf, header_buf...)
	buf = append(buf, question_buf...)
	buf = append(buf, answer_buf...)

	return buf
}


// Parsing functions
func ParseMessage(message_stream []byte) (DNSMessage, error) {
	if len(message_stream) < 12 {
		return DNSMessage{}, fmt.Errorf(`bytestream is fewer than the minimum of 12 bytes for DNS message (got %v bytes).`, len(message_stream))
	}
	parsed_header, err := ParseHeader(message_stream[:12])
	if err != nil {
		return DNSMessage{}, err
	}
	parsed_questions, questions_bytecount, err := ParseQuestions(message_stream[12:], parsed_header.qdcount)
	if err != nil {
		return DNSMessage{}, err
	}
	parsed_answers, _, err := ParseAnswers(message_stream[12+questions_bytecount:], parsed_header.ancount)
	if err != nil {
		return DNSMessage{}, err
	}

	return DNSMessage{header:parsed_header, question:parsed_questions, answer:parsed_answers}, nil
}
