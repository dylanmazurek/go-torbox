package goclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dylanmazurek/go-torbox/pkg/goclient"
	"github.com/stretchr/testify/assert"
)

type TestResponse struct {
	Success bool `json:"success"`
}

func TestDoWithRetry_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client := goclient.NewClient(context.Background(), goclient.WithBaseURL(server.URL))
	req, err := client.NewRequest(http.MethodGet, "/test")
	assert.NoError(t, err)

	var result TestResponse
	err = client.Do(req, &result)
	assert.NoError(t, err)
	assert.Equal(t, true, result.Success)
}

func TestDoWithRetry_RetryOn500(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "server error"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client := goclient.NewClient(context.Background(), goclient.WithBaseURL(server.URL))
	req, err := client.NewRequest(http.MethodGet, "/test")
	assert.NoError(t, err)

	var result TestResponse
	// We expect it to succeed after retries
	err = client.DoWithRetry(req, &result, 3)
	assert.NoError(t, err)
	assert.Equal(t, true, result.Success)
	assert.Equal(t, 3, attempts)
}

func TestDoWithRetry_RetryOn429(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 2 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "too many requests"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	client := goclient.NewClient(context.Background(), goclient.WithBaseURL(server.URL))
	req, err := client.NewRequest(http.MethodGet, "/test")
	assert.NoError(t, err)

	var result TestResponse
	// We expect it to succeed after waiting
	err = client.DoWithRetry(req, &result, 3)
	assert.NoError(t, err)
	assert.Equal(t, true, result.Success)
	assert.Equal(t, 2, attempts)
}

func TestDoWithRetry_FailAfterMaxRetries(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "persistent error"}`))
	}))
	defer server.Close()

	client := goclient.NewClient(context.Background(), goclient.WithBaseURL(server.URL))
	req, err := client.NewRequest(http.MethodGet, "/test")
	assert.NoError(t, err)

	var result TestResponse
	err = client.DoWithRetry(req, &result, 2)
	assert.Error(t, err)
}
