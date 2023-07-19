package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// User prints out the user part from url.
func User(cfg *config.Config) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		parts := parseURL(cfg.URL)

		if parts.User != nil {
			if len(parts.User.Username()) > 0 {
				fmt.Println(parts.User.Username())
			}
		}

		return nil
	}
}
