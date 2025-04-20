package command

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v3"
)

// Run default command and print out full url.
func Run(cfg *config.Config) cli.ActionFunc {
	return func(_ context.Context, _ *cli.Command) error {
		parts := NewURLParser(cfg.URL, cfg.QueryField, cfg.QuerySplit).parse()

		if len(parts.String()) > 0 {
			if cfg.JSONOutput {
				json, _ := json.Marshal(parts)
				fmt.Println(string(json))
			} else {
				fmt.Println(parts)
			}
		}

		return nil
	}
}

// AllFlags defines flags for all subcommand.
func AllFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "json",
			Usage:       "output json",
			Sources:     cli.EnvVars("URL_PARSER_JSON"),
			Destination: &cfg.JSONOutput,
		},
	}
}
