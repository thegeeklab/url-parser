package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// Password prints out the password part from url.
func Password(cfg *config.Config) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		parts := parseURL(cfg.URL)

		if parts.User != nil {
			pw, _ := parts.User.Password()
			if len(pw) > 0 {
				fmt.Println(pw)
			}
		}

		return nil
	}
}
