package promtengineclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type client struct {
	domain     string
	httpClient *http.Client
}

// NewClients creates a client object and check if the information provided are not empty
func NewClient(domain string, httpClient *http.Client) (*client, error) {
	switch {
	case domain == ``:
		return nil, fmt.Errorf(`error_no_domain`)
	case httpClient == nil:
		return nil, fmt.Errorf(`error_no_httpClient`)
	}

	return &client{
		domain:     domain,
		httpClient: &http.Client{},
	}, nil
}

func (c *client) get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.domain, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(`Content-Type`, `application/json`)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}

func (c *client) post(endpoint string, payload string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.domain, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body, err
}
