package command

import (
	"context"
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v3"
)

// Password prints out the password part from url.
func Password(cfg *config.Config) cli.ActionFunc {
	return func(_ context.Context, _ *cli.Command) error {
		parts := NewURLParser(cfg.URL, cfg.QueryField, cfg.QuerySplit).parse()

		if parts.Password != "" {
			fmt.Println(parts.Password)
		}

		return nil
	}
}
