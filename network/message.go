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
		Payload payload.Payload
	}
)

func NewMessage(cmdType CommandType, p payload.Payload) *Message {
	var size uint32
	if p != nil {
		buf := new(bytes.Buffer)
		if err := p.EncodeBinary(buf); err != nil {
			panic(err)
		}
		// TODO: このbufにencodeしたpayloadがある
		size = uint32(buf.Len())
	}

	return &Message{
		Command: cmdToByteArray(cmdType),
		Length:  size,
		Payload: p,
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
	fmt.Println("**** decode")
	if err := binary.Read(r, binary.LittleEndian, &m.Command); err != nil {
		fmt.Println("#### Decode error: Command -> ", err)
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &m.Length); err != nil {
		fmt.Println("#### Decode error: Lenght -> ", err)
		return err
	}
	if m.Length == 0 {
		fmt.Println("Message has no payload")
		return nil
	}

	if err := m.decodePayload(r); err != nil {
		fmt.Println("#### Decode error: Payload -> ", err)
		return err
	}

	return nil
	//return m.decodePayload(r)
}

func (m *Message) decodePayload(r io.Reader) error {
	fmt.Println("decode !!!!!!!!!!!")
	buf := new(bytes.Buffer)
	n, err := io.CopyN(buf, r, int64(m.Length))
	if err != nil {
		fmt.Println("#### Decode error: Payload copy to buffer -> ", err)
		return err
	}

	if uint32(n) != m.Length {
		return fmt.Errorf("expected to doesn't match length, expected: %d, actual: %d", m.Length, n)
	}

	var p payload.Payload
	cmdType := m.CommandType()
	fmt.Println("受信したMessageのcmdType:", cmdType)
	switch cmdType {
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
	fmt.Println("command: ", cmd) // "version"
	switch cmd {
	case "version":
		fmt.Println("バージョンはコマンドだよ, バージョン-> ", cmd)
		return CMDVersion
	case "verack":
		return CMDverack
	default:
		fmt.Println("バージョンはわからないだよ, バージョン-> ", cmd)
		return CMDUnknown
	}
}

func cmdByteArrayToString(byteArr [cmdByte]byte) string {
	buf := []byte{}
	for i := 0; i < cmdByte; i++ {
		if byteArr[i] != 0 {
			buf = append(buf, byteArr[i])
		}
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
