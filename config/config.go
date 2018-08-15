package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type (
	Config struct {
		ProtocolConfiguration    ProtocolConfiguration    `yaml:"ProtocolConfiguration"`
		ApplicationConfiguration ApplicationConfiguration `yaml:"ApplicationConfiguration"`
	}

	ProtocolConfiguration struct {
		SeedList []string `yaml:"SeedList"`
	}

	ApplicationConfiguration struct {
		DataPath    string
		NodePort    uint16
		DialTimeout uint16
		MaxPeers    uint32
	}
)

func Load(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c Config
	if err := yaml.Unmarshal(buf, &c); err != nil {
		return nil, err
	}

	return &c, nil
}
