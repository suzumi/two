package payload

import "io"

type (
	Payload interface {
		EncodeBinary(io.Writer) error
		DecodeBinary(io.Reader) error
	}
)
