package network

import "net"

type (
	Peer interface {
		WriteMsg()
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

func (p *TCPPeer) WriteMsg() {
	//
}

func (p *TCPPeer) Disconnect(err error) {
	p.done <- err
}
