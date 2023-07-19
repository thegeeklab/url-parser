package command

import (
	"strings"
	"testing"

	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

type TestPasswordData struct {
	config   *config.Config
	expected string
}

func TestPassword(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestPasswordData{
		{
			config:   &config.Config{URL: urlString},
			expected: "pass",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Password(table.config)(ctx) }))

		if result != table.expected {
			t.Fatalf("URL password `%v`, should be `%v`", result, table.expected)
		}
	}
}
