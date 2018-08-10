package main

import (
	"github.com/urfave/cli"
	"github.com/suzumi/two/cli/wallet"
	"os"
)

func main() {
	ctl := cli.NewApp()
	ctl.Name = "TWO"
	ctl.Usage = "go client for TWO node"

	ctl.Commands = []cli.Command{
		wallet.NewCommand(),
	}

	ctl.Run(os.Args)
}
