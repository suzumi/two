package wallet

import "github.com/urfave/cli"

func NewCommand() cli.Command {
	return cli.Command{
		Name: "wallet",
		Usage: "create wallet",
		Subcommands: []cli.Command{
			{
				Name: "create",
				Usage: "create a new wallet",
				Action: "createWallet",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name: "path, p",
						Usage: "location of wallet file",
					},
				},
			},
		},
	}
}

func createWallet(ctx *cli.Context) error {
	path := ctx.String("path")
	if len(path) == 0 {
		cli.NewExitError("doesn't specified path wallet file", 1)
	}
	return nil
}