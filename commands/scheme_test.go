package commands

import (
	"flag"
	"strings"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/urfave/cli/v2"
)

type TestSchemeData struct {
	urlString string
	expected  string
}

func TestScheme(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestSchemeData{
		TestSchemeData{
			urlString: urlString,
			expected:  "postgres",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		set := flag.NewFlagSet("test", 0)
		set.String("url", table.urlString, "test url")

		c := cli.NewContext(app, set, nil)
		result := strings.TrimSpace(capturer.CaptureStdout(func() { Scheme(c) }))

		if result != table.expected {
			t.Fatalf("URL scheme `%v`, should be `%v`", result, table.expected)
		}
	}
}
