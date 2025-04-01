package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v3"
)

// Port prints out the port from the url.
func Port(cfg *config.Config) cli.ActionFunc {
	return func(_ *cli.Context) error {
		parts := NewURLParser(cfg.URL, cfg.QueryField, cfg.QuerySplit).parse()

		if len(parts.Scheme) > 0 {
			fmt.Println(parts.Port)
		}

		return nil
	}
}
