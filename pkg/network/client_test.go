package network

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient("https://example.com")
	if client.baseURL != "https://example.com" {
		t.Errorf("Expected baseURL https://example.com, got %s", client.baseURL)
	}

	if client.maxRetries != 3 {
		t.Errorf("Expected maxRetries 3, got %d", client.maxRetries)
	}
}

func TestNewClientWithOptions(t *testing.T) {
	client := NewClient(
		"https://example.com",
		WithTimeout(60*time.Second),
		WithMaxRetries(5),
	)

	if client.maxRetries != 5 {
		t.Errorf("Expected maxRetries 5, got %d", client.maxRetries)
	}
}

func TestClientGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := NewClient(server.URL)
	resp, err := client.Get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var result map[string]string
	if err := json.Unmarshal(resp, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["status"] != "ok" {
		t.Errorf("Expected status ok, got %s", result["status"])
	}
}

func TestClientPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}))
	defer server.Close()

	client := NewClient(server.URL)
	resp, err := client.Post(context.Background(), "/test", map[string]string{"key": "value"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	var result map[string]string
	if err := json.Unmarshal(resp, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["result"] != "success" {
		t.Errorf("Expected result success, got %s", result["result"])
	}
}

func TestClientRetry(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	}))
	defer server.Close()

	client := NewClient(server.URL, WithMaxRetries(3))
	resp, err := client.Get(context.Background(), "/test")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}

	var result map[string]string
	if err := json.Unmarshal(resp, &result); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["status"] != "ok" {
		t.Errorf("Expected status ok, got %s", result["status"])
	}
}

func TestClientError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}))
	defer server.Close()

	client := NewClient(server.URL)
	_, err := client.Get(context.Background(), "/test")
	if err == nil {
		t.Error("Expected error for bad request")
	}
}
