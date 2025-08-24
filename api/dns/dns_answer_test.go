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
