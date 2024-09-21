package infra

import (
	"net/http"
	"time"

	"GoGateway/util"
)

type HTTPClient struct {
	Client *http.Client
	Logger util.Logger
}

func NewHTTPClient(timeout time.Duration, logger util.Logger) *HTTPClient {
	return &HTTPClient{
		Client: &http.Client{
			Timeout: timeout,
		},
		Logger: logger,
	}
}

// Example method to fetch external data
func (hc *HTTPClient) Get(endpoint string) (*http.Response, error) {
	hc.Logger.Info("Making GET request", "endpoint", endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		hc.Logger.Error("Failed to create HTTP request", "error", err)
		return nil, err
	}

	resp, err := hc.Client.Do(req)
	if err != nil {
		hc.Logger.Error("HTTP request failed", "error", err)
		return nil, err
	}

	return resp, nil
}

// Implement other HTTP methods as needed (POST, PUT, DELETE, etc.)
