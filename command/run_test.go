package command

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegeeklab/url-parser/config"
	"github.com/urfave/cli/v3"
	"github.com/zenizh/go-capturer"
)

func TestRun(t *testing.T) {
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tests := []struct {
		name     string
		config   *config.Config
		expected string
	}{
		{
			name:     "get url",
			config:   &config.Config{URL: urlString},
			expected: urlString,
		},
		{
			name: "get url with query split",
			config: &config.Config{
				URL:        urlString,
				QuerySplit: true,
				JSONOutput: true,
			},
			expected: `{
				"scheme": "postgres",
				"hostname": "host.com",
				"port": "5432",
				"path": "/path/to",
				"fragment": "some-fragment",
				"rawQuery": "key=value&other=other%20value",
				"queryParams": [
					{
						"key": "key",
						"value": "value"
					},
					{
						"key": "other",
						"value": "other value"
					}
				],
				"username": "user",
				"password": "pass"
			}`,
		},
	}

	for _, tt := range tests {
		app := cli.NewApp()
		ctx := cli.NewContext(app, nil, nil)

		t.Run(tt.name, func(t *testing.T) {
			result := strings.TrimSpace(capturer.CaptureStdout(func() { _ = Run(tt.config)(ctx) }))

			if tt.config.JSONOutput {
				got := &URL{}
				expected := &URL{}

				_ = json.Unmarshal([]byte(result), &got)
				_ = json.Unmarshal([]byte(tt.expected), &expected)

				assert.Equal(t, expected.Scheme, got.Scheme)
				assert.Equal(t, expected.Username, got.Username)
				assert.Equal(t, expected.Password, got.Password)
				assert.Equal(t, expected.Hostname, got.Hostname)
				assert.Equal(t, expected.Port, got.Port)
				assert.Equal(t, expected.Path, got.Path)
				assert.Equal(t, expected.Fragment, got.Fragment)
				assert.Equal(t, expected.RawQuery, got.RawQuery)
				assert.Equal(t, expected.Query, got.Query)
				assert.ElementsMatch(t, expected.QueryParams, got.QueryParams)

				return
			}

			assert.Equal(t, tt.expected, result)
		})
	}
}
