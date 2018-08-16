package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
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
		DataPath    string        `yaml:"DataPath"`
		NodePort    uint16        `yaml:"NodePort"`
		DialTimeout time.Duration `yaml:"DialTimeout"`
		MaxPeers    uint32        `yaml:"MaxPeers"`
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
