package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mdp/qrterminal"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "oobid"
	app.Version = "1.0.0"
	app.Usage = "generate transfer url for mixin messager"
	app.Commands = commands()

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
	}
}

func flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{Name: "asset"},
		cli.StringFlag{Name: "recipient"},
		cli.StringFlag{Name: "amount"},
		cli.StringFlag{Name: "memo"},
		cli.BoolFlag{Name: "limit"},
		cli.BoolFlag{Name: "market"},
		cli.StringFlag{Name: "price"},
		cli.StringFlag{Name: "target"},
		cli.BoolFlag{Name: "qrcode"},
	}
}

func commands() []cli.Command {
	commands := []cli.Command{}
	for name, action := range map[string]string{"transfer": "", "bid": "B", "ask": "A", "donate": "D"} {
		n, a := name, action
		commands = append(commands, cli.Command{
			Name:  n,
			Flags: flags(),
			Action: func(c *cli.Context) error {
				p, err := extractPayment(c, a)
				if err == nil {
					err = transfer(p, c.Bool("qrcode"))
				}

				return err
			},
		})
	}

	return commands
}

func extractPayment(c *cli.Context, action string) (Payment, error) {
	p := Payment{}

	p.Recipient = c.String("recipient")

	p.Asset = symbolsMap[strings.ToUpper(c.String("asset"))]
	if len(p.Asset) == 0 {
		return p, fmt.Errorf("asset id is not found")
	}

	amount, err := decimal.NewFromString(c.String("amount"))
	if err != nil {
		return p, errors.Wrap(err, "invalid amount")
	}

	if amount.LessThan(decimal.Zero) {
		return p, fmt.Errorf("amount must be positive")
	}

	p.Amount = amount.String()
	p.Memo = c.String("memo")

	if action == "B" || action == "A" {
		p.Recipient = oceanOneId

		priceType := ""

		price, _ := decimal.NewFromString(c.String("price"))
		target := symbolsMap[strings.ToUpper(c.String("target"))]

		if c.Bool("limit") {
			priceType = "L"
		}

		if c.Bool("market") {
			priceType = "M"
			price = decimal.Zero
		}

		if len(priceType) != 0 && len(target) != 0 {
			if memo, err := createMemo(action, price.String(), priceType, target); err != nil {
				log.Println("memo", memo)
				return p, err
			} else {
				p.Memo = memo
			}
		}
	} else if action == "D" {
		p.Recipient = author
	}

	if len(p.Recipient) == 0 {
		return p, fmt.Errorf("recipient is not found")
	}

	return p, nil
}

func transfer(p Payment, qrcode bool) error {
	url := p.PaymentUrl()
	log.Println(url)

	if qrcode {
		qrterminal.Generate(url, qrterminal.L, os.Stdout)
	}

	return nil
}
