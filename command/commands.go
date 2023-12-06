package command

import (
	"net/url"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/thegeeklab/url-parser/config"
)

func parseURL(raw string) *url.URL {
	urlString := strings.TrimSpace(raw)

	url, err := url.Parse(urlString)
	if err != nil {
		log.Fatal().Err(err).Msg(config.ErrParseURL.Error())
	}

	return url
}
