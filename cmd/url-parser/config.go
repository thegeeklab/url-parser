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
			Usage:   "print out all parts from url",
			Action:  commands.Run,
			Flags:   globalFlags(),
		},
		{
			Name:    "scheme",
			Aliases: []string{"s"},
			Usage:   "print out scheme from url",
			Action:  commands.Scheme,
			Flags:   globalFlags(),
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "print out username from url",
			Action:  commands.User,
			Flags:   globalFlags(),
		},
		{
			Name:    "password",
			Aliases: []string{"p"},
			Usage:   "print out password from url",
			Action:  commands.Password,
			Flags:   globalFlags(),
		},
		{
			Name:    "path",
			Aliases: []string{"pt"},
			Usage:   "print out the path from url",
			Action:  commands.Path,
			Flags:   append(globalFlags(), commands.PathFlags()...),
		},
	}
}
