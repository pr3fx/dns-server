package dns

import (
	"testing"
)


var TestVectors = []struct{
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
	for _, testVector := range TestVectors {
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
