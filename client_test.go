package promtengineclient_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	promtengineclient "github.com/BinaryLogos/promt-engine-client"
)

func TestNewClient(t *testing.T) {
	type testCases struct {
		name       string
		domain     string
		httpClient *http.Client
		err        error
	}

	tc := []testCases{
		{
			name:       `empty_domain`,
			domain:     ``,
			httpClient: &http.Client{},
			err:        fmt.Errorf(`error_no_domain`),
		},
		{
			name:       `empty_httpClient`,
			domain:     `https://www.google.com`,
			httpClient: nil,
			err:        fmt.Errorf(`error_no_httpClient`),
		},
		{
			name:       `valid_client`,
			domain:     `https://www.google.com`,
			httpClient: &http.Client{},
			err:        nil,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			_, err := promtengineclient.NewClient(tc.domain, tc.httpClient)
			assert.Equal(t, tc.err, err)
		})
	}
}
