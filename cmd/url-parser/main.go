package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thegeeklab/url-parser/command"
	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v3"
)

//nolint:gochecknoglobals
var (
	BuildVersion = "devel"
	BuildDate    = "00000000"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version=%s date=%s\n", c.App.Name, c.App.Version, BuildDate)
	}

	cfg := &config.Config{}

	app := &cli.App{
		Name:    "url-parser",
		Usage:   "Parse URL and shows the part of it.",
		Version: BuildVersion,
		Action:  command.Run(cfg),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "url",
				Usage:       "source url to parse",
				EnvVars:     []string{"URL_PARSER_URL"},
				Destination: &cfg.URL,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Get all parts from url",
				Action:  command.Run(cfg),
				Flags:   command.AllFlags(cfg),
			},
			{
				Name:    "scheme",
				Aliases: []string{"s"},
				Usage:   "Get scheme from url",
				Action:  command.Scheme(cfg),
			},
			{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "Get username from url",
				Action:  command.User(cfg),
			},
			{
				Name:    "password",
				Aliases: []string{"pw"},
				Usage:   "Get password from url",
				Action:  command.Password(cfg),
			},
			{
				Name:    "path",
				Aliases: []string{"pt"},
				Usage:   "Get path from url",
				Action:  command.Path(cfg),
				Flags:   command.PathFlags(cfg),
			},
			{
				Name:    "host",
				Aliases: []string{"ht"},
				Usage:   "Get hostname from url",
				Action:  command.Host(cfg),
			},
			{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Get port from url",
				Action:  command.Port(cfg),
			},
			{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Get query from url",
				Action:  command.Query(cfg),
				Flags:   command.QueryFlags(cfg),
			},
			{
				Name:    "fragment",
				Aliases: []string{"f"},
				Usage:   "Get fragment from url",
				Action:  command.Fragment(cfg),
			},
		},
		Before: func(_ *cli.Context) error {
			if cfg.URL == "" {
				stat, _ := os.Stdin.Stat()
				if (stat.Mode() & os.ModeCharDevice) == 0 {
					stdin, err := io.ReadAll(os.Stdin)
					if err != nil {
						return fmt.Errorf("error: %w: %w", config.ErrReadStdin, err)
					}
					cfg.URL = strings.TrimSuffix(string(stdin), "\n")
				}
			}

			if cfg.URL == "" {
				return fmt.Errorf("error: %w", config.ErrEmptyURL)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("Execution error")
	}
}
