package command

import (
	"flag"
	"strings"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/urfave/cli/v2"
)

type TestPasswordData struct {
	urlString string
	expected  string
}

func TestPassword(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestPasswordData{
		{
			urlString: urlString,
			expected:  "pass",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		set := flag.NewFlagSet("test", 0)
		set.String("url", table.urlString, "test url")

		c := cli.NewContext(app, set, nil)
		result := strings.TrimSpace(capturer.CaptureStdout(func() { Password(c) }))

		if result != table.expected {
			t.Fatalf("URL password `%v`, should be `%v`", result, table.expected)
		}
	}
}
