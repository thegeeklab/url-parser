package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// QueryFlags defines flags for query subcommand
func QueryFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "query-field",
			Usage:   "filter parsed query string by field name",
			EnvVars: []string{"URL_PARSER_QUERY_FIELD"},
		},
	}
}

// Query prints out the query part from url
func Query(ctx *cli.Context) error {
	parts := parseURL(ctx.String("url"))
	f := ctx.String("query-field")

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
