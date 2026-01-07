package network

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client handles HTTP communication with the Zama relayer.
// Client is safe for concurrent use by multiple goroutines.
type Client struct {
	baseURL    string
	httpClient *http.Client
	maxRetries int
}

// ClientOption configures the Client.
type ClientOption func(*Client)

// WithTimeout sets the HTTP client timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithMaxRetries sets the maximum number of retries.
func WithMaxRetries(retries int) ClientOption {
	return func(c *Client) {
		c.maxRetries = retries
	}
}

// NewClient creates a new network client for the Zama relayer.
func NewClient(baseURL string, opts ...ClientOption) *Client {
	c := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		maxRetries: 3,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Get performs an HTTP GET request.
func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	return c.doRequest(ctx, http.MethodGet, path, nil)
}

// Post performs an HTTP POST request with JSON body.
func (c *Client) Post(ctx context.Context, path string, body interface{}) ([]byte, error) {
	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	return c.doRequest(ctx, http.MethodPost, path, reqBody)
}

func (c *Client) doRequest(ctx context.Context, method, path string, body []byte) ([]byte, error) {
	url := c.baseURL + path

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(attempt*attempt) * 100 * time.Millisecond
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(backoff):
			}
		}

		var bodyReader io.Reader
		if body != nil {
			bodyReader = bytes.NewReader(body)
		}

		req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("request failed: %w", err)
			continue
		}

		respBody, err := io.ReadAll(resp.Body)
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = closeErr
		}

		if err != nil {
			lastErr = fmt.Errorf("failed to read response: %w", err)
			continue
		}

		if resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("server error: %d", resp.StatusCode)
			continue
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
		}

		return respBody, nil
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

// Close closes the client and releases any resources.
// It closes idle connections in the underlying HTTP client.
func (c *Client) Close() error {
	c.httpClient.CloseIdleConnections()
	return nil
}
