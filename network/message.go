package network

import (
	"time"
	"io"
	"encoding/binary"
)

const (
	cmdByte = 12

	CMDVersion CommandType = "version"
	CMDverack  CommandType = "verack"
)

type (
	CommandType string
	Message struct {
		Command [cmdByte]byte
		Length  uint32
		Payload *Payload
	}
)

func NewMessage(cmdType CommandType, payload *Payload) *Message {
	return &Message{
		Command: cmdToByteArray(cmdType),
		Payload: payload,
	}
}

func (m *Message) Encode(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, m.Command); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, m.Length); err != nil {
		return err
	}
	if m.Payload != nil {
		return m.Encode(w)
	}
	return nil
}

func cmdToByteArray(cmd CommandType) [cmdByte]byte {
	cmdLen := len(cmd)
	if cmdLen > cmdByte {
		panic("exceeded max length")
	}

	b := [cmdByte]byte{}
	for i := 0; i < cmdLen; i++ {
		b[i] = cmd[i]
	}
	return b
}

type (
	Payload struct {
		Version     uint32
		Timestamp   uint32
		BlockHeight uint32
		NodeID      uint32
	}
)

func NewPayload(height uint32, id uint32) *Payload {
	return &Payload{
		Version:     0,
		Timestamp:   uint32(time.Now().UTC().Unix()),
		BlockHeight: height,
		NodeID:      id,
	}
}

func (p *Payload) Encode(w io.Writer) error {
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
