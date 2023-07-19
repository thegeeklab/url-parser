package command

import (
	"strings"
	"testing"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

type TestQueryData struct {
	config     *config.Config
	QueryField string
	expected   string
}

func TestQuery(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestQueryData{
		{
			config:   &config.Config{URL: urlString},
			expected: "key=value&other=other%20value",
		},
		{
			config: &config.Config{URL: urlString, QueryField: "other"},

			expected: "other value",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Query(table.config)(ctx) }))

		if result != table.expected {
			t.Fatalf("URL query `%v`, should be `%v`", result, table.expected)
		}
	}
}
