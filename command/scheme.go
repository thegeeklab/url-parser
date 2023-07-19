package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// Scheme prints out the scheme part from the url.
func Scheme(config *config.Config) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		parts := parseURL(config.URL)

		if len(parts.Scheme) > 0 {
			fmt.Println(parts.Scheme)
		}

		return nil
	}
}