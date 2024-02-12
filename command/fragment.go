package command

import (
	"fmt"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
)

// Fragment prints out the fragment part from the url.
func Fragment(cfg *config.Config) cli.ActionFunc {
	return func(_ *cli.Context) error {
		parts := parseURL(cfg.URL)

		if len(parts.Scheme) > 0 {
			fmt.Println(parts.Fragment)
		}

		return nil
	}
}
