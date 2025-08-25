package dns

import (
	"testing"
	"reflect"
)


var TestVectors_QR_OPCODE_AA_TC_RD = []struct{
	qr uint8
	opcode uint8
	aa uint8
	tc uint8
	rd uint8
	want uint8
}{//QR, OPCODE, AA, TC, RD, want
//      1111                01111001
	{0, 15,     0,  0,  1,  121},
//      1010                11010110
	{1, 10,     1,  1,  0,  214},
//      1111                11111111
	{1, 15,     1,  1,  1,  255},
//      0000                10000010
	{1, 0,      0,  1,  0,  130},
}

func TestSet_QR_OPCODE_AA_TC_RD(t *testing.T) {
	header := DNSHeader{}
	var err error
	for _, testVector := range TestVectors_QR_OPCODE_AA_TC_RD {
		err = header.SetQR(testVector.qr)
		if err != nil {
			t.Errorf(`Error when setting QR: %s`, err)
		}
		err = header.SetOPCODE(testVector.opcode)
		if err != nil {
			t.Errorf(`Error when setting OPCODE: %s`, err)
		}
		err = header.SetAA(testVector.aa)
		if err != nil {
			t.Errorf(`Error when setting AA: %s`, err)
		}
		err = header.SetTC(testVector.tc)
		if err != nil {
			t.Errorf(`Error when setting TC: %s`, err)
		}
		err = header.SetRD(testVector.rd)
		if err != nil {
			t.Errorf(`Error when setting RD: %s`, err)
		}
		if header.qr_opcode_aa_tc_rd != testVector.want {
			t.Errorf(`header.qr_opcode_aa_tc_rd = %v, want %v`, header.qr_opcode_aa_tc_rd, testVector.want)
		}
	}
}

var TestVectors_RA_Z_RCODE = []struct{
	ra uint8
	z uint8
	rcode uint8
	want uint8
}{
//   RA, Z,   RCODE, want
//       111  1111   11111111
	{1,  7,   15,    255},
//       010  1000   00101000
	{0,  2,   8,     40},
//       101  1110   11011110
	{1,  5,   14,    222},
//       000  0000   00000000
	{0,  0,   0,     0},
}

func TestSet_RA_Z_RCODE(t *testing.T) {
	header := DNSHeader{}
	var err error
	for _, testVector := range TestVectors_RA_Z_RCODE {
		err = header.SetRA(testVector.ra)
		if err != nil {
			t.Errorf(`Error when setting RA: %s`, err)
		}
		err = header.SetZ(testVector.z)
		if err != nil {
			t.Errorf(`Error when setting Z: %s`, err)
		}
		err = header.SetRCODE(testVector.rcode)
		if err != nil {
			t.Errorf(`Error when setting RCODE: %s`, err)
		}
		if header.ra_z_rcode != testVector.want {
			t.Errorf(`header.ra_z_rcode = %v, want %v`, header.ra_z_rcode, testVector.want)
		}
	}
}

var TestVectors_Serialize = []struct{
	header DNSHeader
	want []byte
}{
	{DNSHeader{65535,255,255,65535,65535,65535,65535}, []byte{255,255,255,255,255,255,255,255,255,255,255,255,}},
	{DNSHeader{0,0,0,0,0,0,0}, []byte{0,0,0,0,0,0,0,0,0,0,0,0,}},
	{DNSHeader{43690,240,29,60331,11179,12267,3}, []byte{170,170,240,29,235,171,43,171,47,235,0,3,}},
}

func TestSerialize(t *testing.T) {
	for _, testVector := range TestVectors_Serialize {
		got_buf := testVector.header.Serialize()
		if !reflect.DeepEqual(got_buf, testVector.want) {
			t.Errorf(`header.Serialize() = %v, want %v`, got_buf, testVector.want)
		}
	}
}


var TestVectors_Parse = []struct{
	header DNSHeader
	want DNSHeader
}{
	{DNSHeader{65535,255,255,65535,65535,65535,65535}, DNSHeader{65535,255,255,65535,65535,65535,65535}},
	{DNSHeader{0,0,0,0,0,0,0}, DNSHeader{0,0,0,0,0,0,0}},
	{DNSHeader{43690,240,29,60331,11179,12267,3}, DNSHeader{43690,240,29,60331,11179,12267,3}},
}

func TestParse(t *testing.T) {
	for _, testVector := range TestVectors_Parse {
		got_header_parsed, err := ParseHeader(testVector.header.Serialize())
		if err != nil {
			t.Errorf(`Encountered error while parsing header: %v`, err)
		}
		if got_header_parsed != testVector.header {
			t.Errorf(`header_parsed = %v want %v`, got_header_parsed, testVector.header)
		}
	}
}
