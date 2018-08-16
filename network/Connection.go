package network

import (
	"net"
	"time"
	"fmt"
)

type Connector struct {
	node     *Node
	listener net.Listener
}

func NewConnector(n *Node) *Connector {
	return &Connector{
		node: n,
	}
}

func (c *Connector) Dial(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return err
	}
	fmt.Println(conn)
	c.connectionHandler(conn)
	return nil
}

func (c *Connector) Accept() {
	//
}

func (c *Connector) connectionHandler(conn net.Conn) {

	var (
		p = NewPeer(conn)
		err error
	)

	defer p.Disconnect(err)

	c.node.Peer <- p

	// decode message

	return
}
