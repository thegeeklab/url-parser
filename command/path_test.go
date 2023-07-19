package command

import (
	"strings"
	"testing"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

type TestPathData struct {
	config   *config.Config
	expected string
}

func TestPath(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestPathData{
		{
			config:   &config.Config{URL: urlString, PathIndex: -1},
			expected: "/path/to",
		},
		{
			config:   &config.Config{URL: urlString, PathIndex: 0},
			expected: "path",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Path(table.config)(ctx) }))

		if result != table.expected {
			t.Fatalf("URL path `%v`, should be `%v`", result, table.expected)
		}
	}
}
