package command

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v3"
	"github.com/zenizh/go-capturer"
)

func TestPath(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tests := []struct {
		name     string
		config   *config.Config
		expected string
	}{
		{
			name:     "get path",
			config:   &config.Config{URL: urlString, PathIndex: -1},
			expected: "/path/to",
		},
		{
			name:     "get path at index",
			config:   &config.Config{URL: urlString, PathIndex: 0},
			expected: "path",
		},
	}

	for _, tt := range tests {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		t.Run(tt.name, func(t *testing.T) {
			result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Path(tt.config)(ctx) }))
			assert.Equal(t, tt.expected, result)
		})
	}
}
