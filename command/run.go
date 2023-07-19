package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// Run default command and print out full url.
func Run(cfg *config.Config) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		parts := parseURL(cfg.URL)

		if len(parts.String()) > 0 {
			fmt.Println(parts)
		}

		return nil
	}
}
