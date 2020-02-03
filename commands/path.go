package commands

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

// PathFlags defines flags for path subcommand
func PathFlags() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:    "path-index",
			Usage:   "filter parsed path by index",
			EnvVars: []string{"URL_PARSER_PATH_INDEX"},
			Value:   -1,
		},
	}
}

// Path prints out the path part from url
func Path(ctx *cli.Context) error {
	parts := parseURL(ctx.String("url"))
	i := ctx.Int("path-index")

	if len(parts.Path) > 0 {
		if i > -1 {
			path := strings.Split(parts.Path, "/")

			if i = i + 1; i < len(path) {
				fmt.Println(path[i])
			}
		} else {
			fmt.Println(parts.Path)
		}
	}

	return nil
}
