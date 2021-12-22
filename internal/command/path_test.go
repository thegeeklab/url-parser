package command

import (
	"flag"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

type TestPathData struct {
	urlString string
	pathIndex int
	expected  string
}

func TestPath(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestPathData{
		{
			urlString: urlString,
			pathIndex: -1,
			expected:  "/path/to",
		},
		{
			urlString: urlString,
			pathIndex: 0,
			expected:  "path",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		set := flag.NewFlagSet("test", 0)
		set.String("url", table.urlString, "test url")
		set.Int("path-index", table.pathIndex, "index")

		c := cli.NewContext(app, set, nil)
		result := strings.TrimSpace(capturer.CaptureStdout(func() { Path(c) }))

		if result != table.expected {
			t.Fatalf("URL path `%v`, should be `%v`", result, table.expected)
		}
	}
}
