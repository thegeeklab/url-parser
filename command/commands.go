package command

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/url-parser/config"
)

func parseURL(raw string) *url.URL {
	urlString := strings.TrimSpace(raw)

	url, err := url.Parse(urlString)
	if err != nil {
		logrus.Fatal(fmt.Errorf("%w: %w", config.ErrParseURL, err))
	}

	return url
}
