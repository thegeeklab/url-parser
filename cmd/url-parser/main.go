package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/xoxys/url-parser/commands"
)

// Version of current build
var Version = "devel"

func main() {
	app := cli.NewApp()
	app.Name = "url-parser"
	app.Usage = "Parse URL and shows the part of it."
	app.Version = Version
	app.Action = commands.Run
	app.Flags = globalFlags()
	app.Commands = configCommands()

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
