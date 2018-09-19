package main

import (
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	app := configureCommandLineParsing()
	app.Run(os.Args)
}

func configureCommandLineParsing() *cli.App {
	app := cli.NewApp()
	app.Name = "shekel"
	app.Usage = "A command-line client for OAuth 2.0"

	app.Commands = []cli.Command{
		{
			Name:   "oauth2",
			Usage:  "stuff",
			Action: executeOAuth2Request,
		},
	}

	return app
}

func executeOAuth2Request(c *cli.Context) error {
	fmt.Println("oauth stuff!")
	return nil
}
