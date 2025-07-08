package dns


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
