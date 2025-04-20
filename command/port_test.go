package command

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v3"
	"github.com/zenizh/go-capturer"
)

func TestPort(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tests := []struct {
		name     string
		config   *config.Config
		expected string
	}{
		{
			name:     "get port",
			config:   &config.Config{URL: urlString},
			expected: "5432",
		},
	}

	for _, tt := range tests {
		app := &cli.Command{}

		t.Run(tt.name, func(t *testing.T) {
			result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Port(tt.config)(t.Context(), app) }))
			assert.Equal(t, tt.expected, result)
		})
	}
}
