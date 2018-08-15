package network

type (
	NodeConfig struct {
		SeedList      []string
		ListenTCPPort uint16
		DialTimeout   uint16
		MaxPeers      uint16
	}

	Node struct {
		NodeConfig
		ID uint32
		Peer Peer
	}

	Peer struct {
		//
	}
)
