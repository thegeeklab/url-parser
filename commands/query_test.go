package commands

import (
	"flag"
	"strings"
	"testing"

	"github.com/kami-zh/go-capturer"
	"github.com/urfave/cli/v2"
)

type TestQueryData struct {
	urlString  string
	QueryField string
	expected   string
}

func TestQuery(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestQueryData{
		{
			urlString: urlString,
			expected:  "key=value&other=other%20value",
		},
		{
			urlString:  urlString,
			QueryField: "other",
			expected:   "other value",
		},
	}

	for _, table := range tables {
		app := cli.NewApp()
		set := flag.NewFlagSet("test", 0)
		set.String("url", table.urlString, "test url")
		set.String("query-field", table.QueryField, "index")

		c := cli.NewContext(app, set, nil)
		result := strings.TrimSpace(capturer.CaptureStdout(func() { Query(c) }))

		if result != table.expected {
			t.Fatalf("URL query `%v`, should be `%v`", result, table.expected)
		}
	}
}
