package network

import (
	"net"
	"fmt"
)

type (
	Peer interface {
		WriteMsg(*Message) error
		Disconnect(error)
	}

	TCPPeer struct {
		conn net.Conn
		done chan error
	}
)

func NewPeer(conn net.Conn) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		done: make(chan error),
	}
}

func (p *TCPPeer) WriteMsg(msg *Message) error {
	fmt.Println("send version messages....")
	return msg.Encode(p.conn)
}

func (p *TCPPeer) Disconnect(err error) {
	p.done <- err
}
