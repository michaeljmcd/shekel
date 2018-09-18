package main

import (
        "os"
        "fmt"
        "github.com/urfave/cli"
       )

func main() {
    app := configureCommandLineParsing()
    app.Run(os.Args)

    //fmt.Println("Hello")
}

func configureCommandLineParsing() *cli.App {
//askCommand := flag.NewFlagSet("oauth2", flag.ExitOnError)
//    askCommand.String("stuff", "", "Stuff")
    app := cli.NewApp()
    app.Commands = []cli.Command{
        {
            Name: "oauth2",
            Usage: "stuff",
            Action: func(c *cli.Context) error {
                fmt.Println("oauth stuff!")
                return nil
            },
        },
    }

    return app
}
