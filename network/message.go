package network

const cmdByte = 12

type (
	Message struct {
		Command [cmdByte]byte
		Length  uint32
		Payload Payload
	}

	Payload struct {
		//
	}
)
