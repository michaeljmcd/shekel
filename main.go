package main

import (
    "context"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
    "log"
	"os"
)

var (
    logger = log.New(os.Stderr, "logger: ", log.Lshortfile)
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
			Usage:  "Execute commands using OAuth 2.0 flows.",
            Flags: []cli.Flag{
                cli.StringFlag {
                    Name: "auth-url",
                    Value: "http://example.com/authurl",
                    Usage: "Authorization URL for your OAuth 2.0 provider",
                },
                cli.StringFlag {
                    Name: "token-url",
                    Value: "http://example.com/tokenurl",
                    Usage: "Token URL for your OAuth 2.0 provider",
                },
                cli.StringFlag {
                    Name: "code",
                    Value: "",
                    Usage: "An authorization code retrieved from a sign-in screen.",
                },
                cli.StringFlag {
                    Name: "resource",
                    Value: "resource",
                    Usage: "A resource to be requested using the provided information.",
                },
                cli.StringFlag {
                    Name: "client-id",
                    Value: "Client ID",
                    Usage: "A client ID to use.",
                },
                cli.StringFlag {
                    Name: "client-secret",
                    Value: "Client Secret.",
                    Usage: "Client secret to be used when gaining access.",
                },
            },
			Action: executeOAuth2Request,
		},
	}

	return app
}

func executeOAuth2Request(c *cli.Context) error {
    logger.Print("Creating a client")

    ctx := context.Background()
    conf := &oauth2.Config {
        ClientID: c.String("client-id"),
        ClientSecret: c.String("client-secret"),
        Scopes: []string{"SCOPE1"},
        Endpoint: oauth2.Endpoint {
            AuthURL: c.String("auth-url"),
            TokenURL: c.String("token-url"),
        },
    }

    authorizationCode := c.String("code")

    if authorizationCode == "" {
        url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
        logger.Printf("Visit the url %v for an authorization code.", url)
    } else {
        token, err := conf.Exchange(ctx, authorizationCode)

        if err != nil {
            logger.Fatal(err)
        }

        client := conf.Client(ctx, token)
        client.Get("")
    }

	return nil
}
