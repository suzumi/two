package wallet

import (
	"github.com/urfave/cli"
	"bufio"
	"os"
	"fmt"
	"strings"
	"errors"
	"github.com/suzumi/two/wallet"
)

func NewCommand() cli.Command {
	return cli.Command{
		Name:  "wallet",
		Usage: "create wallet",
		Subcommands: []cli.Command{
			{
				Name:   "create",
				Usage:  "create a new wallet",
				Action: createWallet,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "path, p",
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

	wlt, err := wallet.NewWallet(path)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	if err := wlt.Save(); err != nil {
		return cli.NewExitError(err, 1)
	}

	if err := createAccount(ctx, wlt); err != nil {
		return cli.NewExitError(err, 1)
	}

	dumpWallet(wlt)

	return nil
}

func createAccount(ctx *cli.Context, wlt *wallet.Wallet) error {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Enter wallet pass phrase > ")
	rawPhrase, _ := buf.ReadBytes('\n')
	fmt.Print("Confirm wallet pass phrase > ")
	rawPhraseCheck, _ := buf.ReadBytes('\n')

	phrase := strings.TrimRight(string(rawPhrase), "\n")
	phraseCheck := strings.TrimRight(string(rawPhraseCheck), "\n")

	if phrase != phraseCheck {
		return errors.New("Entered pass phrase does not match")
	}


	return wlt.CreateAccount(phrase)
}

func dumpWallet(wlt *wallet.Wallet)  {
	b, _ := wlt.JSON()
	fmt.Println("")
	fmt.Println(string(b))
	fmt.Println("")
}
