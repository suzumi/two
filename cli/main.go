package main

import (
	"github.com/urfave/cli"
	"github.com/suzumi/two/cli/wallet"
	"os"
	"github.com/suzumi/two/cli/server"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	ctl := cli.NewApp()
	ctl.Name = "TWO"
	ctl.Usage = "go client for TWO node"

	ctl.Commands = []cli.Command{
		wallet.NewCommand(),
		server.NewCommand(),
	}

	ctl.Run(os.Args)
}
