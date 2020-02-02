package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// User prints out the user part from url
func User(ctx *cli.Context) error {
	parts := parseURL(ctx)

	if parts.User != nil {
		if len(parts.User.Username()) > 0 {
			fmt.Println(parts.User.Username())
		}
	}
	return nil
}
