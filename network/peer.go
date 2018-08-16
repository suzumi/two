package network

import "net"

type Peer struct {
	conn net.Conn
	done chan error
}

func NewPeer(conn net.Conn) *Peer {
	return &Peer{
		conn: conn,
		done: make(chan error),
	}
}

func (p *Peer) WriteMsg() {
	//
}

func (p *Peer) Disconnect(err error) {
	p.done <- err
}
