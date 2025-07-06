package dns

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"encoding/binary"
)


type DNSHeader struct {
	id uint16                 // Packet Identifier

                              // {1 bit}         {4 bits}        {1 bit}               {1 bit}            {1 bit}
	qr_opcode_aa_tc_rd uint8  // Query Response, Operation Code, Authoritative Answer, Truncated Message, Recursion Desired

                              // {1 bit}              {3 bits}  {4 bits}
	ra_z_rcode uint8          // Recursion Available, Reserved, Response Code

	qdcount uint16            // Question Count
	ancount uint16            // Answer Count
	nscount uint16            // Authority Count
	arcount uint16            // Additional Count
}


// DNSHeader field-setting functions
func (header *DNSHeader) SetID(ID uint16) {
	(*header).id = ID
}

func (header *DNSHeader) SetQR(QR uint8) {
	if QR > 1 {
		log.Warning(fmt.Sprintf("Attempting to set QR bit with a value greater than 1 (%v)", QR))
		log.Warning("Only the first bit from the input (QR) will be used, others will be ignored.")
	}
	(*header).qr_opcode_aa_tc_rd |= QR & uint8(1) << 7
}

func (header *DNSHeader) SetOPCODE(OPCODE uint8) {
	// Take only 4 LSBs
	(*header).qr_opcode_aa_tc_rd |= OPCODE & uint8(15) << 3
}

func (header *DNSHeader) SetAA(AA uint8) {
	if AA > 1 {
		log.Warning(fmt.Sprintf("Attempting to set AA bit with a value greater than 1 (%v)", AA))
		log.Warning("Only the first bit from the input (AA) will be used, others will be ignored.")
	}
	(*header).qr_opcode_aa_tc_rd |= AA & uint8(1) << 2
}

func (header *DNSHeader) SetTC(TC uint8) {
	if TC > 1 {
		log.Warning(fmt.Sprintf("Attempting to set TC bit with a value greater than 1 (%v)", TC))
		log.Warning("Only the first bit from the input (TC) will be used, others will be ignored.")
	}
	(*header).qr_opcode_aa_tc_rd |= TC & uint8(1) << 1
}

func (header *DNSHeader) SetRD(RD uint8) {
	if RD > 1 {
		log.Warning(fmt.Sprintf("Attempting to set RD bit with a value greater than 1 (%v)", RD))
		log.Warning("Only the first bit from the input (RD) will be used, others will be ignored.")
	}
	(*header).qr_opcode_aa_tc_rd |= RD & uint8(1)
}

func (header *DNSHeader) SetRA(RA uint8) {
	if RA > 1 {
		log.Warning(fmt.Sprintf("Attempting to set RA bit with a value greater than 1 (%v)", RA))
		log.Warning("Only the first bit from the input (RA) will be used, others will be ignored.")
	}
	(*header).ra_z_rcode |= RA & uint8(1) << 7
}

func (header *DNSHeader) SetZ(Z uint8) {
	// Take only 3 LSBs
	(*header).ra_z_rcode |= Z & uint8(7) << 4
}

func (header *DNSHeader) SetRCODE(RCODE uint8) {
	// Take only 4 LSBs
	(*header).ra_z_rcode |= RCODE & uint8(15)
}

func (header *DNSHeader) SetQDCOUNT(QDCOUNT uint16) {
	(*header).qdcount = QDCOUNT
}

func (header *DNSHeader) SetANCOUNT(ANCOUNT uint16) {
	(*header).ancount = ANCOUNT
}

func (header *DNSHeader) SetNSCOUNT(NSCOUNT uint16) {
	(*header).nscount = NSCOUNT
}

func (header *DNSHeader) SetARCOUNT(ARCOUNT uint16) {
	(*header).arcount = ARCOUNT
}

// Return a byte slice of the struct
func (header DNSHeader) Serialize() []byte {
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[0:], header.id)
	buf[2] = header.qr_opcode_aa_tc_rd
	buf[3] = header.ra_z_rcode
	binary.BigEndian.PutUint16(buf[4:], header.qdcount)
	binary.BigEndian.PutUint16(buf[6:], header.ancount)
	binary.BigEndian.PutUint16(buf[8:], header.nscount)
	binary.BigEndian.PutUint16(buf[10:], header.arcount)
	return buf
}

// Print header contents
func (header DNSHeader) PrintFields() {
	fmt.Println("---------------------DNS HEADER---------------------")
	fmt.Println("RFC Name    Bitfield            Descriptive Name")
	fmt.Println("----------------------------------------------------")
	fmt.Printf( "ID          %016b    Packet Identifier\n", header.id)
	fmt.Printf( "QR          %01b                   Query Response\n", header.qr_opcode_aa_tc_rd >> 7)
	fmt.Printf( "OPCODE      %04b                Operation Code\n", header.qr_opcode_aa_tc_rd >> 3 & uint8(15))
	fmt.Printf( "AA          %01b                   Authoritative Answer\n", header.qr_opcode_aa_tc_rd >> 2 & uint8(1))
	fmt.Printf( "TC          %01b                   Truncated Message\n", header.qr_opcode_aa_tc_rd >> 1 & uint8(1))
	fmt.Printf( "RD          %01b                   Recursion Desired\n", header.qr_opcode_aa_tc_rd & uint8(1))
	fmt.Printf( "RA          %01b                   Recursion Available\n", header.ra_z_rcode >> 7)
	fmt.Printf( "Z           %03b                 Reserved\n", header.ra_z_rcode >> 4 & uint8(7))
	fmt.Printf( "RCODE       %04b                Response Code\n", header.ra_z_rcode & uint8(15))
	fmt.Printf( "QDCOUNT     %016b    Question Count\n", header.qdcount)
	fmt.Printf( "ANCOUNT     %016b    Answer Count\n", header.ancount)
	fmt.Printf( "NSCOUNT     %016b    Authority Count\n", header.nscount)
	fmt.Printf( "ARCOUNT     %016b    Additional Count\n", header.arcount)
	fmt.Println("----------------------------------------------------")
}
