package command

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v2"
	"github.com/zenizh/go-capturer"
)

func TestPassword(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tests := []struct {
		name     string
		config   *config.Config
		expected string
	}{
		{
			name:     "get password",
			config:   &config.Config{URL: urlString},
			expected: "pass",
		},
	}

	for _, tt := range tests {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		t.Run(tt.name, func(t *testing.T) {
			result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Password(tt.config)(ctx) }))
			assert.Equal(t, tt.expected, result)
		})
	}
}
