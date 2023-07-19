package command

import (
	"fmt"
	"strings"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// PathFlags defines flags for path subcommand.
func PathFlags(config *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:        "path-index",
			Usage:       "filter parsed path by index",
			EnvVars:     []string{"URL_PARSER_PATH_INDEX"},
			Value:       -1,
			Destination: &config.PathIndex,
		},
	}
}

// Path prints out the path part from url.
func Path(config *config.Config) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		parts := parseURL(config.URL)
		i := config.PathIndex

		if len(parts.Path) > 0 {
			if i > -1 {
				path := strings.Split(parts.Path, "/")

				if i++; i < len(path) {
					fmt.Println(path[i])
				}
			} else {
				fmt.Println(parts.Path)
			}
		}

		return nil
	}
}
