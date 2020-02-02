package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/xoxys/url-parser/commands"
)

var (
	version = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "url-parser"
	app.Usage = "Parse URL and shows the part of it."
	app.Version = version
	app.Action = commands.Run
	app.Flags = globalFlags()
	app.Commands = configCommands()

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
