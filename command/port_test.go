package command

import (
	"strings"
	"testing"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

type TestPortData struct {
	config   *config.Config
	expected string
}

func TestPort(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestPortData{
		{
			config:   &config.Config{URL: urlString},
			expected: "5432",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Port(table.config)(ctx) }))

		if result != table.expected {
			t.Fatalf("URL port `%v`, should be `%v`", result, table.expected)
		}
	}
}
