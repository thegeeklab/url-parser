package main

import (
	"github.com/urfave/cli/v2"
	"github.com/xoxys/url-parser/commands"
)

func globalFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "url",
			Usage:   "source url to parse",
			EnvVars: []string{"URL_PARSER_URL"},
		},
	}
}

func configCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "Get all parts from url",
			Action:  commands.Run,
			Flags:   globalFlags(),
		},
		{
			Name:    "scheme",
			Aliases: []string{"s"},
			Usage:   "Get scheme from url",
			Action:  commands.Scheme,
			Flags:   globalFlags(),
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Get username from url",
			Action:  commands.User,
			Flags:   globalFlags(),
		},
		{
			Name:    "password",
			Aliases: []string{"pw"},
			Usage:   "Get password from url",
			Action:  commands.Password,
			Flags:   globalFlags(),
		},
		{
			Name:    "path",
			Aliases: []string{"pt"},
			Usage:   "Get path from url",
			Action:  commands.Path,
			Flags:   append(globalFlags(), commands.PathFlags()...),
		},
		{
			Name:    "host",
			Aliases: []string{"h"},
			Usage:   "Get hostname from url",
			Action:  commands.Host,
			Flags:   globalFlags(),
		},
		{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "Get port from url",
			Action:  commands.Port,
			Flags:   globalFlags(),
		},
		{
			Name:    "query",
			Aliases: []string{"q"},
			Usage:   "Get query from url",
			Action:  commands.Query,
			Flags:   append(globalFlags(), commands.QueryFlags()...),
		},
		{
			Name:    "fragment",
			Aliases: []string{"f"},
			Usage:   "Get fragment from url",
			Action:  commands.Fragment,
			Flags:   globalFlags(),
		},
	}
}
