package main

import (
	"github.com/thegeeklab/url-parser/internal/command"
	"github.com/urfave/cli/v2"
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
			Action:  command.Run,
			Flags:   globalFlags(),
		},
		{
			Name:    "scheme",
			Aliases: []string{"s"},
			Usage:   "Get scheme from url",
			Action:  command.Scheme,
			Flags:   globalFlags(),
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Get username from url",
			Action:  command.User,
			Flags:   globalFlags(),
		},
		{
			Name:    "password",
			Aliases: []string{"pw"},
			Usage:   "Get password from url",
			Action:  command.Password,
			Flags:   globalFlags(),
		},
		{
			Name:    "path",
			Aliases: []string{"pt"},
			Usage:   "Get path from url",
			Action:  command.Path,
			Flags:   append(globalFlags(), command.PathFlags()...),
		},
		{
			Name:    "host",
			Aliases: []string{"h"},
			Usage:   "Get hostname from url",
			Action:  command.Host,
			Flags:   globalFlags(),
		},
		{
			Name:    "port",
			Aliases: []string{"p"},
			Usage:   "Get port from url",
			Action:  command.Port,
			Flags:   globalFlags(),
		},
		{
			Name:    "query",
			Aliases: []string{"q"},
			Usage:   "Get query from url",
			Action:  command.Query,
			Flags:   append(globalFlags(), command.QueryFlags()...),
		},
		{
			Name:    "fragment",
			Aliases: []string{"f"},
			Usage:   "Get fragment from url",
			Action:  command.Fragment,
			Flags:   globalFlags(),
		},
	}
}
