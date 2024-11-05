package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegeeklab/url-parser/config"
)

func TestParse(t *testing.T) {
	//nolint:goconst
	urlString := "postgres://user:pass@host.com:5432/path/to?key=value&other=other%20value#some-fragment"

	tests := []struct {
		name     string
		config   *config.Config
		expected *URL
	}{
		{
			name: "parse url",
			config: &config.Config{
				URL:        urlString,
				QuerySplit: true,
			},
			expected: &URL{
				Scheme:   "postgres",
				Username: "user",
				Password: "pass",
				Hostname: "host.com",
				Port:     "5432",
				Path:     "/path/to",
				Query:    "key=value&other=other%20value",
				RawQuery: "key=value&other=other%20value",
				QueryParams: []QueryParam{
					{
						Key:   "key",
						Value: "value",
					},
					{
						Key:   "other",
						Value: "other value",
					},
				},
				Fragment: "some-fragment",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewURLParser(urlString, "", false).parse()
			assert.Equal(t, tt.expected.Scheme, result.Scheme)
			assert.Equal(t, tt.expected.Username, result.Username)
			assert.Equal(t, tt.expected.Password, result.Password)
			assert.Equal(t, tt.expected.Hostname, result.Hostname)
			assert.Equal(t, tt.expected.Port, result.Port)
			assert.Equal(t, tt.expected.Path, result.Path)
			assert.Equal(t, tt.expected.Fragment, result.Fragment)
			assert.Equal(t, tt.expected.RawQuery, result.RawQuery)
			assert.Equal(t, tt.expected.Query, result.Query)
			assert.ElementsMatch(t, tt.expected.QueryParams, result.QueryParams)
		})
	}
}
