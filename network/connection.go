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

const TCP = "tcp"

func NewConnector(n *Node) *Connector {
	return &Connector{
		node: n,
	}
}

func (c *Connector) Dial(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout(TCP, addr, timeout*time.Second)
	if err != nil {
		fmt.Println("call Dial error ...")
		return err
	}
	go c.connectionHandler(conn)
	return nil
}

func (c *Connector) Accept() {
	l, err := net.Listen(TCP, fmt.Sprintf(":%d", c.node.ApplicationConfiguration.NodePort))
	if err != nil {
		fmt.Printf("listen TCP error: %s\n", err)
		return
	}

	c.listener = l

	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		go c.connectionHandler(conn)
	}
}

func (c *Connector) connectionHandler(conn net.Conn) {

	var (
		p   = NewPeer(conn)
		err error
	)

	fmt.Printf("★Connection handler, remote node: %s\n", conn.RemoteAddr())

	defer p.Disconnect(err)

	c.node.Register <- p

	for {
		msg := &Message{}
		if err := msg.Decode(p.conn); err != nil {
			fmt.Printf("Decode Error: %e\n", err)
			return
		}
		fmt.Printf("**Message: %s\n", msg.Command)
	}

}
