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
	conn, err := net.DialTimeout(TCP, addr, timeout)
	if err != nil {
		return err
	}
	fmt.Println(conn)
	c.connectionHandler(conn)
	return nil
}

func (c *Connector) Accept() {
	fmt.Printf("listen to port: %d\n", c.node.ApplicationConfiguration.NodePort)
	l, err := net.Listen(TCP, fmt.Sprintf(":%d", c.node.ApplicationConfiguration.NodePort))
	if err != nil {
		fmt.Printf("listen TCP error: %s\n", err)
		return
	}

	c.listener = l

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("")
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

	fmt.Println("connection handler")

	defer p.Disconnect(err)

	c.node.Peer <- p

	// decode message

}
