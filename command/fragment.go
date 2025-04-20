package command

import (
	"context"
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v3"
)

// Fragment prints out the fragment part from the url.
func Fragment(cfg *config.Config) cli.ActionFunc {
	return func(_ context.Context, _ *cli.Command) error {
		parts := NewURLParser(cfg.URL, cfg.QueryField, cfg.QuerySplit).parse()

		if len(parts.Scheme) > 0 {
			fmt.Println(parts.Fragment)
		}

		return nil
	}
}
