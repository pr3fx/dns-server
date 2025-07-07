package dns


type DNSMessage struct {
	header DNSHeader
	question DNSQuestion
}

func (message *DNSMessage) SetHeader(header DNSHeader) {
	(*message).header = header
}

func (message *DNSMessage) SetQuestion(question DNSQuestion) {
	(*message).question = question
}
