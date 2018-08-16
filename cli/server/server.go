package server

import (
	"github.com/urfave/cli"
	"fmt"
	"github.com/suzumi/two/config"
	"github.com/suzumi/two/network"
)

func NewCommand() cli.Command {
	return cli.Command{
		Name:   "node",
		Usage:  "start a TWO node",
		Action: startServer,
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "mainnet, m"},
		},
	}
}

func startServer(ctx *cli.Context) error {
	net := "testnet"
	if ctx.Bool("mainnet") {
		net = "mainnet"
	}
	configPath := fmt.Sprintf("./config/%s.yml", net)
	conf, err := config.Load(configPath)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	fmt.Println(conf)

	node := network.NewNode(conf)

	done := make(chan bool)
	node.Start()
	<-done

	return nil
}
