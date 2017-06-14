package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/wonsikin/icotjo/parser"
)

const (
	// AppVersion the version of this app.
	AppVersion = "v0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Version = AppVersion
	app.Name = "icotjo"
	app.HelpName = "icotjo"
	app.Usage = "a utility for parse i18n.csv to i18n json file"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Value: "./i18N.csv",
			Usage: "where is the i18n.csv file",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "./",
			Usage: "the output of the swagger file that was generated",
		},
	}

	app.Action = func(c *cli.Context) error {
		err := parser.Parser(c.String("input"), c.String("output"))
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("[Error] %v", err)
	}
}
