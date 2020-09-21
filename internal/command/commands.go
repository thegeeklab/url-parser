package command

import (
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

func parseURL(raw string) *url.URL {
	urlString := strings.TrimSpace(raw)

	url, err := url.Parse(urlString)
	if err != nil {
		logrus.Fatal(err)
	}

	return url
}
