package commands

import (
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func parseURL(c *cli.Context) *url.URL {
	urlString := strings.TrimSpace(c.String("url"))

	url, err := url.Parse(urlString)
	if err != nil {
		logrus.Fatal(err)
	}

	return url
}
