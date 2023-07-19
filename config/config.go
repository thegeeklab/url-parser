package config

import "errors"

var (
	ErrRequiredFlagsNotSet = errors.New("either \"url\" or \"stdin\" must be set")
	ErrExclusiveFlags      = errors.New("\"url\" and \"stdin\" are mutually exclusive")
	ErrEmptyStdin          = errors.New("\"stdin\" must not be empty")
	ErrReadStdin           = errors.New("failed to read \"stdin\"")
	ErrParseURL            = errors.New("failed to parse url")
)

type Config struct {
	URL        string
	Stdin      bool
	QueryField string
	PathIndex  int
}
