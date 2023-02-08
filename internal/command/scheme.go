package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Scheme prints out the scheme part from the url.
func Scheme(ctx *cli.Context) error {
	parts := parseURL(ctx.String("url"))

	if len(parts.Scheme) > 0 {
		fmt.Println(parts.Scheme)
	}

	return nil
}
