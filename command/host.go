package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// Host prints out the host part from the url.
func Host(cfg *config.Config) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		parts := parseURL(cfg.URL)

		if len(parts.Scheme) > 0 {
			fmt.Println(parts.Hostname())
		}

		return nil
	}
}
