package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Host prints out the host part from the url
func Host(ctx *cli.Context) error {
	parts := parseURL(ctx.String("url"))

	if len(parts.Scheme) > 0 {
		fmt.Println(parts.Hostname())
	}
	return nil
}
