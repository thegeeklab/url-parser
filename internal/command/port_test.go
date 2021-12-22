package command

import (
	"flag"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

type TestPortData struct {
	urlString string
	expected  string
}

func TestPort(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestPortData{
		{
			urlString: urlString,
			expected:  "5432",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		set := flag.NewFlagSet("test", 0)
		set.String("url", table.urlString, "test url")

		c := cli.NewContext(app, set, nil)
		result := strings.TrimSpace(capturer.CaptureStdout(func() { Port(c) }))

		if result != table.expected {
			t.Fatalf("URL port `%v`, should be `%v`", result, table.expected)
		}
	}
}
