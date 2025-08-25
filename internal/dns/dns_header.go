package dns

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"encoding/binary"
	"errors"
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

func (header *DNSHeader) SetQR(QR uint8) error {
	if QR > 1 {
		log.Error(fmt.Sprintf("Attempting to set QR bit with a value greater than 1 (%v)", QR))
		return errors.New("Input QR value greater than 1")
	}
	mask := uint8(128) // 10000000
	if QR == 1 {
		(*header).qr_opcode_aa_tc_rd |= mask
	} else {
		(*header).qr_opcode_aa_tc_rd &= ^mask
	}
	return nil
}

func (header *DNSHeader) SetOPCODE(OPCODE uint8) error {
	if OPCODE > 15 {
		log.Error(fmt.Sprintf("Attempting to set OPCODE with a value greater than 15 (%v)", OPCODE))
		return errors.New("Input OPCODE value greater than 15")
	}
	// Clear OPCODE field
	mask := uint8(135) // 10000111
	(*header).qr_opcode_aa_tc_rd &= mask
	// Set OPCODE bits
	(*header).qr_opcode_aa_tc_rd |= OPCODE << 3
	return nil
}

func (header *DNSHeader) SetAA(AA uint8) error {
	if AA > 1 {
		log.Error(fmt.Sprintf("Attempting to set AA bit with a value greater than 1 (%v)", AA))
		return errors.New("Input AA value greater than 1")
	}
	mask := uint8(4) // 00000100
	if AA == 1 {
		(*header).qr_opcode_aa_tc_rd |= mask
	} else {
		(*header).qr_opcode_aa_tc_rd &= ^mask
	}
	return nil
}

func (header *DNSHeader) SetTC(TC uint8) error {
	if TC > 1 {
		log.Error(fmt.Sprintf("Attempting to set TC bit with a value greater than 1 (%v)", TC))
		return errors.New("Input TC value greater than 1")
	}
	mask := uint8(2) // 00000010
	if TC == 1 {
		(*header).qr_opcode_aa_tc_rd |= mask
	} else {
		(*header).qr_opcode_aa_tc_rd &= ^mask
	}
	return nil
}

func (header *DNSHeader) SetRD(RD uint8) error {
	if RD > 1 {
		log.Error(fmt.Sprintf("Attempting to set RD bit with a value greater than 1 (%v)", RD))
		return errors.New("Input RD value greater than 1")
	}
	mask := uint8(1) // 00000001
	if RD == 1 {
		(*header).qr_opcode_aa_tc_rd |= mask
	} else {
		(*header).qr_opcode_aa_tc_rd &= ^mask
	}
	return nil
}

func (header *DNSHeader) SetRA(RA uint8) error {
	if RA > 1 {
		log.Error(fmt.Sprintf("Attempting to set RA bit with a value greater than 1 (%v)", RA))
		return errors.New("Input RA value greater than 1")
	}
	mask := uint8(128) // 10000000
	if RA == 1 {
		(*header).ra_z_rcode |= mask
	} else {
		(*header).ra_z_rcode &= ^mask
	}
	return nil
}

func (header *DNSHeader) SetZ(Z uint8) error {
	if Z > 7 {
		log.Error(fmt.Sprintf("Attempting to set Z bit with a value greater than 7 (%v)", Z))
		return errors.New("Input Z value greater than 7")
	}
	// Clear Z bits
	mask := uint8(143) // 10001111
	(*header).ra_z_rcode &= mask
	// Set Z bits
	(*header).ra_z_rcode |= Z << 4
	return nil
}

func (header *DNSHeader) SetRCODE(RCODE uint8) error {
	if RCODE > 15 {
		log.Error(fmt.Sprintf("Attempting to set RCODE bit with a value greater than 15 (%v)", RCODE))
		return errors.New("Input RCODE value greater than 15")
	}
	// Clear RCODE bits
	mask := uint8(240) // 11110000
	(*header).ra_z_rcode &= mask
	// Set RCODE bits
	(*header).ra_z_rcode |= RCODE
	return nil
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
