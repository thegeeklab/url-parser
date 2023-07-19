package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// Port prints out the port from the url.
func Port(config *config.Config) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		parts := parseURL(config.URL)

		if len(parts.Scheme) > 0 {
			fmt.Println(parts.Port())
		}

		return nil
	}
}
