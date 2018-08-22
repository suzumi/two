package network

import (
	"io"
	"encoding/binary"
	"github.com/suzumi/two/network/payload"
	"bytes"
	"fmt"
)

const (
	cmdByte = 12

	CMDVersion CommandType = "version"
	CMDverack  CommandType = "verack"
	CMDUnknown CommandType = "unknown"
)

type (
	CommandType string
	Message struct {
		// command is 12 byte
		Command [cmdByte]byte

		// length of payload
		Length uint32

		// payload is message
		Payload *payload.Version
	}
)

func NewMessage(cmdType CommandType, payload *payload.Version) *Message {
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
		return m.Payload.EncodeBinary(w)
	}
	return nil
}

func (m *Message) Decode(r io.Reader) error {
	fmt.Println("################ decode")
	if err := binary.Read(r, binary.LittleEndian, m.Command); err != nil {
		fmt.Println("+++ decode error: Command")
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, m.Length); err != nil {
		fmt.Println("+++ decode error: Length")
		return err
	}
	if m.Length == 0 {
		return nil
	}

	if err := m.decodePayload(r); err != nil {
		return err
	}
	fmt.Printf("#### Message: %s", *m)
	return nil
	//return m.decodePayload(r)
}

func (m *Message) decodePayload(r io.Reader) error {
	buf := new(bytes.Buffer)
	n, err := io.Copy(buf, r)
	if err != nil {
		return err
	}

	if uint32(n) != m.Length {
		fmt.Errorf("expected to doesn't match length, expected: %d, actual: %d", m.Length, n)
	}

	var p payload.Payload
	switch m.CommandType() {
	case CMDVersion:
		p = &payload.Version{}
		if err := p.DecodeBinary(r); err != nil {
			return err
		}
		return nil
	case CMDverack:
		return nil
	}
	return nil
}

func (m *Message) CommandType() CommandType {
	cmd := cmdByteArrayToString(m.Command)
	switch cmd {
	case "version":
		return CMDVersion
	case "verack":
		return CMDverack
	default:
		return CMDUnknown
	}
}

func cmdByteArrayToString(byteArr [cmdByte]byte) string {
	buf := []byte{}
	for i := 0; i < cmdByte; i++ {
		buf = append(buf, byteArr[i])
	}
	return string(buf)
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
