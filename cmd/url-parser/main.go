package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/url-parser/command"
	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

//nolint:gochecknoglobals
var (
	BuildVersion = "devel"
	BuildDate    = "00000000"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version=%s date=%s\n", c.App.Name, c.App.Version, BuildDate)
	}

	config := &config.Config{}

	app := &cli.App{
		Name:    "url-parser",
		Usage:   "Parse URL and shows the part of it.",
		Version: BuildVersion,
		Action:  command.Run(config),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "url",
				Usage:       "source url to parse",
				EnvVars:     []string{"URL_PARSER_URL"},
				Destination: &config.URL,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Get all parts from url",
				Action:  command.Run(config),
			},
			{
				Name:    "scheme",
				Aliases: []string{"s"},
				Usage:   "Get scheme from url",
				Action:  command.Scheme(config),
			},
			{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "Get username from url",
				Action:  command.User(config),
			},
			{
				Name:    "password",
				Aliases: []string{"pw"},
				Usage:   "Get password from url",
				Action:  command.Password(config),
			},
			{
				Name:    "path",
				Aliases: []string{"pt"},
				Usage:   "Get path from url",
				Action:  command.Path(config),
				Flags:   command.PathFlags(config),
			},
			{
				Name:    "host",
				Aliases: []string{"h"},
				Usage:   "Get hostname from url",
				Action:  command.Host(config),
			},
			{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Get port from url",
				Action:  command.Port(config),
			},
			{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Get query from url",
				Action:  command.Query(config),
				Flags:   command.QueryFlags(config),
			},
			{
				Name:    "fragment",
				Aliases: []string{"f"},
				Usage:   "Get fragment from url",
				Action:  command.Fragment(config),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
