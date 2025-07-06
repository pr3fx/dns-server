package dns


type DNSHeader struct {
	ID uint16                 // Packet Identifier

                              // {1 bit}         {4 bits}        {1 bit}               {1 bit}           {1 bit}
	QR_OPCODE_AA_TC_RD uint8  // Query Response, Operation Code, Authoritative Answer, Trucated Message, Recursion Desired

                              // {1 bit}              {3 bits}  {4 bits}
	RA_Z_RCODE uint8          // Recursion Available, Reserved, Response Code

	QDCOUNT uint16            // Question Count
	ANCOUNT uint16            // Answer Count
	NSCOUNT uint16            // Authority Count
	ARCOUNT uint16            // Additional Count
}
