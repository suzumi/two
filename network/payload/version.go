package payload

import (
	"time"
	"encoding/binary"
	"io"
)

type (
	Version struct {
		Version     uint32
		Timestamp   uint32
		BlockHeight uint32
		NodeID      uint32
	}
)

func NewVersion(height uint32, id uint32) *Version {
	return &Version{
		Version:     0,
		Timestamp:   uint32(time.Now().UTC().Unix()),
		BlockHeight: height,
		NodeID:      id,
	}
}

func (p *Version) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, p.Version); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, p.Timestamp); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, p.BlockHeight); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, p.NodeID); err != nil {
		return err
	}

	return nil
}

func (p *Version) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, p.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, p.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, p.BlockHeight); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, p.NodeID); err != nil {
		return err
	}
	return nil
}
