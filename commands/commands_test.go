package commands

import "testing"

type TestParseData struct {
	urlString string
	expected  string
}

func TestParseURL(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tables := []TestParseData{
		{
			urlString: urlString,
			expected:  urlString,
		},
	}

	for _, table := range tables {
		result := parseURL(urlString)

		if result.String() != table.expected {
			t.Fatalf("URL `%v`, should be `%v`", result, table.expected)
		}
	}
}
