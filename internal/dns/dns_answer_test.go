package dns

import (
	"testing"
	"net"
	"reflect"
)


func TestAnswerConstructor(t *testing.T) {
	answer_type, err := NewTypeA_Answer(net.IPv4(8,8,8,8))
	if err != nil {
		t.Errorf(`Error encountered when creating new answer_type: %v`, err)
	}
	answer := NewDNSAnswer("codecrafters.io", uint16(1), uint32(155), answer_type)
	if !reflect.DeepEqual(answer.name, []byte{12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0}) {
		t.Errorf(`answer.name = %v, want %v`, answer.name, []byte{12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0})
	}
	if answer.ans_type != RecordType_A {
		t.Errorf(`answer.ans_type = %v, want %v`, answer.ans_type, RecordType_A)
	}
	if answer.class != uint16(1) {
		t.Errorf(`answer.class = %v, want %v`, answer.class, uint16(1))
	}
	if answer.ttl != uint32(155) {
		t.Errorf(`answer.ttl = %v, want %v`, answer.ttl, uint32(155))
	}
	if answer.rdlength != uint32(4) {
	  	t.Errorf(`answer.dlength = %v, want %v`, answer.rdlength, uint32(4))
	}
	if !reflect.DeepEqual(answer.rdata, []byte{8,8,8,8}) {
		t.Errorf(`answer.rdata = %v, want %v`, answer.rdata, []byte{8,8,8,8})
	}
}

var TestVectors_ParseAnswers = []struct{
	ancount uint16
	answer_bytestream []byte
	want_answer_list []DNSAnswer
	want_answer_bytecount int
}{
	{
		uint16(2),
		append(
			DNSAnswer{[]byte{3,119,119,119,2,117,112,2,97,99,2,122,97,0},RecordType_A,uint16(1),uint32(4294967295),uint32(0),[]byte{}}.Serialize(),
			DNSAnswer{[]byte{3,119,119,119,2,117,112,2,97,99,2,122,97,0},RecordType_A,uint16(1),uint32(4294901760),uint32(4),[]byte{10,20,30,40}}.Serialize()...
		),
		[]DNSAnswer{
			DNSAnswer{[]byte{3,119,119,119,2,117,112,2,97,99,2,122,97,0},RecordType_A,uint16(1),uint32(4294967295),uint32(0),[]byte{}},
			DNSAnswer{[]byte{3,119,119,119,2,117,112,2,97,99,2,122,97,0},RecordType_A,uint16(1),uint32(4294901760),uint32(4),[]byte{10,20,30,40}},
		},
		56,
	},
}

func TestParseAnswers(t *testing.T) {
	for _, testVector := range TestVectors_ParseAnswers {
		got_answer_list, got_answer_bytecount, err := ParseAnswers(testVector.answer_bytestream, testVector.ancount)
		if err != nil {
			t.Errorf(`Error encountered while parsing answers: %v`, err)
		}
		if !reflect.DeepEqual(got_answer_list, testVector.want_answer_list) {
			t.Errorf(`got_answer_list = %v, want %v`, got_answer_list, testVector.want_answer_list)
		}
		if got_answer_bytecount != testVector.want_answer_bytecount {
			t.Errorf(`got_answer_bytecount = %v, want %v`, got_answer_bytecount, testVector.want_answer_bytecount)
		}
	}
}
