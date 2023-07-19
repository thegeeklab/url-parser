package command

import (
	"strings"
	"testing"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

type TestSchemeData struct {
	config   *config.Config
	expected string
}

func TestScheme(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestSchemeData{
		{
			config:   &config.Config{URL: urlString},
			expected: "postgres",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Scheme(table.config)(ctx) }))

		if result != table.expected {
			t.Fatalf("URL scheme `%v`, should be `%v`", result, table.expected)
		}
	}
}
