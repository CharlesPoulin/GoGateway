package ports

import "net/http"

type HTTPClient interface {
	Get(endpoint string) (*http.Response, error)
	// Define other HTTP methods as needed
}
