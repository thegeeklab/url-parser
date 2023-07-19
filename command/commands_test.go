package command

import (
	"testing"

	"github.com/thegeeklab/url-parser/config"
)

type TestParseData struct {
	config   *config.Config
	expected string
}

func TestParseURL(t *testing.T) {
	//nolint:goconst
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestParseData{
		{
			config:   &config.Config{URL: urlString},
			expected: urlString,
		},
	}

	for _, table := range tables {
		result := parseURL(urlString)

		if result.String() != table.expected {
			t.Fatalf("URL `%v`, should be `%v`", result, table.expected)
		}
	}
}
