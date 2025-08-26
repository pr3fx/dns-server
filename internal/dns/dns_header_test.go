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


type testHeaderGettersWant struct {
	id uint16
	qr uint8
	opcode uint8
	aa uint8
	tc uint8
	rd uint8
	ra uint8
	z uint8
	rcode uint8
	qdcount uint16
	ancount uint16
	nscount uint16
	arcount uint16
}

var TestVectors_HeaderGetters = []struct{
	header DNSHeader
	want testHeaderGettersWant
}{                    //        1  15     1  1  1  1  7 15
	{                 // id     qr_opcode_aa_tc_rd ra_z_rcode qdcount ancount nscount arcount
		header:DNSHeader{65535, 255,               255,       0,      0,      10,     0},
		want:testHeaderGettersWant{
			id:65535,
			qr:1,
			opcode:15,
			aa:1,
			tc:1,
			rd:1,
			ra:1,
			z:7,
			rcode:15,
			qdcount:0,
			ancount:0,
			nscount:10,
			arcount:0,
		},
	},
                      //        0  10     0  1  0  0  2 11
	{                 // id     qr_opcode_aa_tc_rd ra_z_rcode qdcount     ancount     nscount arcount
		header:DNSHeader{22,    82,                43,        65535,      65500,      2,      55},
		want:testHeaderGettersWant{
			id:22,
			qr:0,
			opcode:10,
			aa:0,
			tc:1,
			rd:0,
			ra:0,
			z:2,
			rcode:11,
			qdcount:65535,
			ancount:65500,
			nscount:2,
			arcount:55,
		},
	},
}

func TestHeaderGetters(t *testing.T) {
	for _, testVector := range TestVectors_HeaderGetters {
		got_id := testVector.header.GetID()
		if got_id != testVector.want.id {
			t.Errorf(`got_id = %v, want = %v`, got_id, testVector.want.id)
		}
		got_qr := testVector.header.GetQR()
		if got_qr != testVector.want.qr {
			t.Errorf(`got_qr = %v, want = %v`, got_qr, testVector.want.qr)
		}
		got_opcode := testVector.header.GetOPCODE()
		if got_opcode != testVector.want.opcode {
			t.Errorf(`got_opcode = %v, want = %v`, got_opcode, testVector.want.opcode)
		}
		got_aa := testVector.header.GetAA()
		if got_aa != testVector.want.aa {
			t.Errorf(`got_aa = %v, want = %v`, got_aa, testVector.want.aa)
		}
		got_tc := testVector.header.GetTC()
		if got_tc != testVector.want.tc {
			t.Errorf(`got_tc = %v, want = %v`, got_tc, testVector.want.tc)
		}
		got_rd := testVector.header.GetRD()
		if got_rd != testVector.want.rd {
			t.Errorf(`got_rd = %v, want = %v`, got_rd, testVector.want.rd)
		}
		got_ra := testVector.header.GetRA()
		if got_ra != testVector.want.ra {
			t.Errorf(`got_ra = %v, want = %v`, got_ra, testVector.want.ra)
		}
		got_z := testVector.header.GetZ()
		if got_z != testVector.want.z {
			t.Errorf(`got_z = %v, want = %v`, got_z, testVector.want.z)
		}
		got_rcode := testVector.header.GetRCODE()
		if got_rcode != testVector.want.rcode {
			t.Errorf(`got_rcode = %v, want = %v`, got_rcode, testVector.want.rcode)
		}
		got_qdcount := testVector.header.GetQDCOUNT()
		if got_qdcount != testVector.want.qdcount {
			t.Errorf(`got_qdcount = %v, want = %v`, got_qdcount, testVector.want.qdcount)
		}
		got_ancount := testVector.header.GetANCOUNT()
		if got_ancount != testVector.want.ancount {
			t.Errorf(`got_ancount = %v, want = %v`, got_ancount, testVector.want.ancount)
		}
		got_nscount := testVector.header.GetNSCOUNT()
		if got_nscount != testVector.want.nscount {
			t.Errorf(`got_nscount = %v, want = %v`, got_nscount, testVector.want.nscount)
		}
		got_arcount := testVector.header.GetARCOUNT()
		if got_arcount != testVector.want.arcount {
			t.Errorf(`got_arcount = %v, want = %v`, got_arcount, testVector.want.arcount)
		}
	}
}
