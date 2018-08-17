package network

import (
	"github.com/suzumi/two/util"
	"github.com/suzumi/two/config"
	"fmt"
)

type (
	Node struct {
		config.Config
		ID        uint32
		Peer      chan Peer
		Connector *Connector
	}
)

func NewNode(config *config.Config) *Node {
	node := &Node{
		Config: *config,
		ID:     util.RandUint32(1000000, 9999999),
		Peer:   make(chan Peer),
	}
	node.Connector = NewConnector(node)
	return node
}

func (n *Node) Start() {
	fmt.Println("node started...")
	go n.ConnectToPeers()
	go n.ListenToTCP()
	n.run()
}

func (n *Node) ConnectToPeers() {
	for _, addr := range n.ProtocolConfiguration.SeedList {
		if err := n.Connector.Dial(addr, n.ApplicationConfiguration.DialTimeout); err != nil {
			fmt.Sprintf("failed to connect: %s\n", err)
		}
	}
	fmt.Println("connecting to peers...")
}

func (n *Node) ListenToTCP() {
	fmt.Println("listen to TCP from peers...")
	n.Connector.Accept()
}

func (n *Node) run() {
	for {
		select {
		case p := <-n.Peer:
			fmt.Println("new peer connected", p)
		}
	}
}
