package command

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

func TestQuery(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tests := []struct {
		name       string
		config     *config.Config
		QueryField string
		expected   string
	}{
		{
			name:     "get query",
			config:   &config.Config{URL: urlString},
			expected: "key=value&other=other%20value",
		},
		{
			name:     "get query field",
			config:   &config.Config{URL: urlString, QueryField: "other"},
			expected: "other value",
		},
	}

	for _, tt := range tests {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		t.Run(tt.name, func(t *testing.T) {
			result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Query(tt.config)(ctx) }))
			assert.Equal(t, tt.expected, result)
		})
	}
}
