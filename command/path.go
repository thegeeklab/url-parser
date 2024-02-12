package command

import (
	"fmt"
	"strings"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// PathFlags defines flags for path subcommand.
func PathFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:        "path-index",
			Usage:       "filter parsed path by index",
			EnvVars:     []string{"URL_PARSER_PATH_INDEX"},
			Value:       -1,
			Destination: &cfg.PathIndex,
		},
	}
}

// Path prints out the path part from url.
func Path(cfg *config.Config) cli.ActionFunc {
	return func(_ *cli.Context) error {
		parts := parseURL(cfg.URL)
		i := cfg.PathIndex

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
