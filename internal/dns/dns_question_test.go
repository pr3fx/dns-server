package dns

import (
	"testing"
	"reflect"
)


var TestVectors_QNAME = []struct{
	qname string
	want []byte
}{
	{"codecrafters.io", []byte{12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0}},
	{"www.codecrafters.io", []byte{3,119,119,119,12,99,111,100,101,99,114,97,102,116,101,114,115,2,105,111,0}},
	{"www.up.ac.za", []byte{3,119,119,119,2,117,112,2,97,99,2,122,97,0}},
}

func TestSetQNAME(t *testing.T) {
	for _, testVector := range TestVectors_QNAME {
		question := DNSQuestion{}
		question.SetQNAME(testVector.qname)
		if !reflect.DeepEqual(question.qname, testVector.want) {
			t.Errorf(`question.qname = %v, want %v`, question.qname, testVector.want)
		}
	}
}
