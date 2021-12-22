package command

import (
	"flag"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

type TestHostnameData struct {
	urlString string
	expected  string
}

func TestHost(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestHostnameData{
		{
			urlString: urlString,
			expected:  "host.com",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		set := flag.NewFlagSet("test", 0)
		set.String("url", table.urlString, "test url")

		c := cli.NewContext(app, set, nil)
		result := strings.TrimSpace(capturer.CaptureStdout(func() { Host(c) }))

		if result != table.expected {
			t.Fatalf("URL host `%v`, should be `%v`", result, table.expected)
		}
	}
}
