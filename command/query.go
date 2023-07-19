package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// QueryFlags defines flags for query subcommand.
func QueryFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "query-field",
			Usage:       "filter parsed query string by field name",
			EnvVars:     []string{"URL_PARSER_QUERY_FIELD"},
			Destination: &cfg.QueryField,
		},
	}
}

// Query prints out the query part from url.
func Query(cfg *config.Config) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		parts := parseURL(cfg.URL)
		f := cfg.QueryField

		if len(parts.RawQuery) > 0 {
			if f != "" {
				if result := parts.Query().Get(f); result != "" {
					fmt.Println(result)
				}
			} else {
				fmt.Println(parts.RawQuery)
			}
		}

		return nil
	}
}
