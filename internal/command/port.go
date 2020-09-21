package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Port prints out the port from the url
func Port(ctx *cli.Context) error {
	parts := parseURL(ctx.String("url"))

	if len(parts.Scheme) > 0 {
		fmt.Println(parts.Port())
	}
	return nil
}
