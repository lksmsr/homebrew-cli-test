package api

import (
	"io"
	"net/http"
	"sync"
)

var (
	clientInstance *SliplaneApiClient
	once           sync.Once
)

// Init initializes the singleton API client. Only the first call has effect.
func Init(apiKey, organizationID string) {
	once.Do(func() {
		clientInstance = NewClient(apiKey, organizationID)
	})
}

// GetClient returns the singleton API client instance, or nil if not initialized.
func GetClient() *SliplaneApiClient {
	return clientInstance
}

type SliplaneApiClient struct {
	APIKey         string
	OrganizationID string
	BaseURL        string
	HTTPClient     *http.Client
}

func NewClient(apiKey, organizationID string) *SliplaneApiClient {
	return &SliplaneApiClient{
		APIKey:         apiKey,
		OrganizationID: organizationID,
		BaseURL:        "https://ctrl.sliplane.io/v0/",
		HTTPClient:     &http.Client{},
	}
}

func (c *SliplaneApiClient) newRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIKey)
	req.Header.Set("X-Organization-ID", c.OrganizationID)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *SliplaneApiClient) Get(path string) (*http.Response, error) {
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}

func (c *SliplaneApiClient) Post(path string, body io.Reader) (*http.Response, error) {
	req, err := c.newRequest(http.MethodPost, path, body)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}

func (c *SliplaneApiClient) Put(path string, body io.Reader) (*http.Response, error) {
	req, err := c.newRequest(http.MethodPut, path, body)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}

func (c *SliplaneApiClient) Patch(path string, body io.Reader) (*http.Response, error) {
	req, err := c.newRequest(http.MethodPatch, path, body)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}

func (c *SliplaneApiClient) Delete(path string) (*http.Response, error) {
	req, err := c.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	return c.HTTPClient.Do(req)
}
