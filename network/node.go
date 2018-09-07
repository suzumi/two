package network

import (
	"github.com/suzumi/two/util"
	"github.com/suzumi/two/config"
	"fmt"
	"github.com/suzumi/two/network/payload"
)

type (
	Node struct {
		config.Config
		ID        uint32
		Register  chan Peer
		Connector *Connector
	}
)

func NewNode(config *config.Config) *Node {
	node := &Node{
		Config:   *config,
		ID:       util.RandUint32(1000000, 9999999),
		Register: make(chan Peer),
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
		fmt.Printf("connecting to peers to ...: %s\n", addr)
		if err := n.Connector.Dial(addr, n.ApplicationConfiguration.DialTimeout); err != nil {
			fmt.Sprintf("failed to connect: %s\n", err)
		}
	}
}

func (n *Node) ListenToTCP() {
	fmt.Println("listen to TCP from peers...")
	n.Connector.Accept()
}

func (n *Node) run() {
	for {
		select {
		case p := <-n.Register:
			fmt.Println("new peer connected", p)
			// send version
			n.sendVersion(p)
		}
	}
}

func (n *Node) sendVersion(p Peer) error {
	// TODO: fix
	pl := payload.NewVersion(50, n.ID)
	newMsg := NewMessage(CMDVersion, pl)
	fmt.Println("new message: ", *newMsg)
	return p.WriteMsg(newMsg)
}
