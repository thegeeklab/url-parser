package commands

import (
	"flag"
	"strings"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/urfave/cli/v2"
)

type TestFragmentData struct {
	urlString string
	expected  string
}

func TestFragment(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestFragmentData{
		{
			urlString: urlString,
			expected:  "some-fragment",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		set := flag.NewFlagSet("test", 0)
		set.String("url", table.urlString, "test url")

		c := cli.NewContext(app, set, nil)
		result := strings.TrimSpace(capturer.CaptureStdout(func() { Fragment(c) }))

		if result != table.expected {
			t.Fatalf("URL fragment `%v`, should be `%v`", result, table.expected)
		}
	}
}
