package command

import (
	"net/url"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/thegeeklab/url-parser/config"
)

type QueryParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type URL struct {
	url *url.URL

	Scheme      string       `json:"scheme"`
	Hostname    string       `json:"hostname"`
	Port        string       `json:"port"`
	Path        string       `json:"path"`
	Fragment    string       `json:"fragment"`
	RawQuery    string       `json:"rawQuery"`
	Query       string       `json:"-"`
	QueryParams []QueryParam `json:"queryParams"`
	Username    string       `json:"username"`
	Password    string       `json:"password"`
}

func (u *URL) String() string {
	return u.url.String()
}

type Parser struct {
	URL        string
	QueryField string
	QuerySplit bool
}

func NewURLParser(url, queryField string, querySplit bool) *Parser {
	return &Parser{
		URL:        url,
		QueryField: queryField,
		QuerySplit: querySplit,
	}
}

func (p *Parser) parse() *URL {
	urlString := strings.TrimSpace(p.URL)

	parts, err := url.Parse(urlString)
	if err != nil {
		log.Fatal().Err(err).Msg(config.ErrParseURL.Error())
	}

	extURL := &URL{
		url:         parts,
		Scheme:      parts.Scheme,
		Hostname:    parts.Hostname(),
		Path:        parts.Path,
		Fragment:    parts.Fragment,
		QueryParams: []QueryParam{},
	}

	if len(parts.Scheme) > 0 {
		extURL.Hostname = parts.Hostname()
		extURL.Port = parts.Port()
	}

	if parts.User != nil {
		if len(parts.User.Username()) > 0 {
			extURL.Username = parts.User.Username()
		}
	}

	if parts.User != nil {
		pw, _ := parts.User.Password()
		if len(pw) > 0 {
			extURL.Password = pw
		}
	}

	// Handle query field extraction
	if parts.RawQuery != "" {
		extURL.RawQuery = parts.RawQuery
	}

	if p.QueryField != "" {
		if result := parts.Query().Get(p.QueryField); result != "" {
			extURL.Query = result
		}
	} else {
		extURL.Query = parts.RawQuery
	}

	// Handle query parameter splitting
	values := parts.Query()
	for k, v := range values {
		if len(v) > 0 {
			extURL.QueryParams = append(extURL.QueryParams, QueryParam{
				Key:   k,
				Value: v[0],
			})
		}
	}

	return extURL
}
