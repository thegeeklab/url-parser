package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// Scheme prints out the scheme part from the url.
func Scheme(cfg *config.Config) cli.ActionFunc {
	return func(_ *cli.Context) error {
		parts := NewURLParser(cfg.URL, cfg.QueryField, cfg.QuerySplit).parse()

		if len(parts.Scheme) > 0 {
			fmt.Println(parts.Scheme)
		}

		return nil
	}
}
