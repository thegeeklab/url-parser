package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Password prints out the password part from url
func Password(ctx *cli.Context) error {
	parts := parseURL(ctx.String("url"))

	if parts.User != nil {
		pw, _ := parts.User.Password()
		if len(pw) > 0 {
			fmt.Println(pw)
		}
	}

	return nil
}
