package config

import "errors"

var (
	ErrEmptyURL  = errors.New("no url provided either by \"url\" or \"stdin\"")
	ErrReadStdin = errors.New("failed to read \"stdin\"")
	ErrParseURL  = errors.New("failed to parse url")
)

type Config struct {
	URL        string
	QueryField string
	QuerySplit bool
	PathIndex  int
	JSONOutput bool
}
