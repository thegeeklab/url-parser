package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// User prints out the user part from url.
func User(cfg *config.Config) cli.ActionFunc {
	return func(_ *cli.Context) error {
		parts := NewURLParser(cfg.URL, cfg.QueryField, cfg.QuerySplit).parse()

		if parts.Username != "" {
			fmt.Println(parts.Username)
		}

		return nil
	}
}
