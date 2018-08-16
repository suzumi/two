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
}

func (n *Node) ConnectToPeers() {
	for _, addr := range n.ProtocolConfiguration.SeedList {
		n.Connector.Dial(addr, n.ApplicationConfiguration.DialTimeout)
	}
	fmt.Println("connecting to peers...")
}

func (n *Node) ListenToTCP() {
	fmt.Println("listen to TCP from peers...")
}
