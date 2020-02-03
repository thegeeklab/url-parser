package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Fragment prints out the fragment part from the url
func Fragment(ctx *cli.Context) error {
	parts := parseURL(ctx.String("url"))

	if len(parts.Scheme) > 0 {
		fmt.Println(parts.Fragment)
	}
	return nil
}
