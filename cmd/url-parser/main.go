package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/url-parser/internal/command"
	"github.com/urfave/cli/v2"
)

var (
	BuildVersion = "devel"
	BuildDate    = "00000000"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version=%s date=%s\n", c.App.Name, c.App.Version, BuildDate)
	}

	app := cli.NewApp()
	app.Name = "url-parser"
	app.Usage = "Parse URL and shows the part of it."
	app.Version = BuildVersion
	app.Action = command.Run
	app.Flags = globalFlags()
	app.Commands = configCommands()

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
