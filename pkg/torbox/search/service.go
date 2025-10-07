package search

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	form "github.com/dylanmazurek/go-torbox/internal/form"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/perimeterx/marshmallow"
	"github.com/rs/zerolog/log"
)

type SearchService struct {
	BaseURL string

	internalClient *http.Client
}

func New(internalClient http.Client) *SearchService {
	return &SearchService{
		BaseURL: constants.API_SEARCH_BASE_URL,

		internalClient: &internalClient,
	}
}

func (s *SearchService) newRequest(method string, reqPath string, urlParams *url.Values, body any) (*http.Request, error) {
	path := fmt.Sprintf("%s/%s", s.BaseURL, reqPath)
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		bodyBuffer, contentType, err := form.ParseMultipartForm(body)
		if err != nil {
			return nil, err
		}

		req.Body = io.NopCloser(bodyBuffer)
		req.Header.Set("Content-Type", contentType)
	}

	if urlParams != nil {
		req.URL.RawQuery = urlParams.Encode()
	}

	return req, nil
}

func (s *SearchService) do(req *http.Request, obj any) error {
	return s.doWithRetry(req, obj, 3)
}

func (s *SearchService) doWithRetry(req *http.Request, obj any, maxRetries int) error {
	var lastErr error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			// Calculate backoff delay: 1s, 2s, 4s
			delay := time.Duration(1<<(attempt-1)) * time.Second
			log.Debug().
				Int("attempt", attempt).
				Dur("delay", delay).
				Msg("retrying torbox search API request after delay")
			time.Sleep(delay)
		}

		req.Header.Add("Accept", "application/json")

		httpResponse, err := s.internalClient.Do(req)
		if err != nil {
			lastErr = err
			log.Warn().
				Err(err).
				Int("attempt", attempt).
				Msg("torbox search API request failed")

			if attempt < maxRetries && isRetryableNetworkError(err) {
				continue
			}

			log.Error().Err(err).Msg("failed to execute torbox search request after retries")
			return err
		}

		if httpResponse != nil && httpResponse.StatusCode >= 400 {
			bodyBytes, _ := io.ReadAll(httpResponse.Body)
			httpResponse.Body.Close()

			log.Debug().
				Str("status", httpResponse.Status).
				Str("message", string(bodyBytes)).
				Int("attempt", attempt).
				Msg("torbox search API response error")

			// Handle rate limiting (429) with Retry-After header
			if httpResponse.StatusCode == 429 {
				retryAfter := httpResponse.Header.Get("Retry-After")
				if retryAfter != "" && attempt < maxRetries {
					if seconds, parseErr := strconv.Atoi(retryAfter); parseErr == nil {
						retryDelay := time.Duration(seconds) * time.Second
						log.Warn().
							Dur("retry_after", retryDelay).
							Int("attempt", attempt).
							Msg("rate limited by torbox search API, waiting before retry")
						time.Sleep(retryDelay)
						continue
					}
				}

				if attempt < maxRetries {
					// Fallback to exponential backoff if no Retry-After header
					delay := time.Duration(1<<attempt) * time.Second
					log.Warn().
						Dur("delay", delay).
						Int("attempt", attempt).
						Msg("rate limited by torbox search API, using exponential backoff")
					time.Sleep(delay)
					continue
				}
			}

			// Handle 5xx server errors with retry
			if httpResponse.StatusCode >= 500 && httpResponse.StatusCode < 600 && attempt < maxRetries {
				log.Warn().
					Int("status_code", httpResponse.StatusCode).
					Int("attempt", attempt).
					Msg("server error from torbox search API, retrying")
				continue
			}

			lastErr = fmt.Errorf("http request failed with status: %s", httpResponse.Status)
			return lastErr
		}

		defer httpResponse.Body.Close()

		if httpResponse.ContentLength != 0 {
			bodyBytes, err := io.ReadAll(httpResponse.Body)
			if err != nil {
				lastErr = err
				if attempt < maxRetries {
					log.Warn().
						Err(err).
						Int("attempt", attempt).
						Msg("failed to read torbox search response body, retrying")
					continue
				}
				return err
			}

			unknownFields, err := marshmallow.Unmarshal(bodyBytes, obj, marshmallow.WithExcludeKnownFieldsFromMap(true))
			if err != nil {
				lastErr = err
				if attempt < maxRetries {
					log.Warn().
						Err(err).
						Int("attempt", attempt).
						Msg("failed to unmarshal torbox search response, retrying")
					continue
				}
				return err
			}

			if len(unknownFields) > 0 {
				log.Warn().Fields(unknownFields).
					Msgf("unknown fields in torbox search response")
			}
		}

		// Success - return without error
		return nil
	}

	return lastErr
}

func isRetryableNetworkError(err error) bool {
	if err == nil {
		return false
	}

	errStr := err.Error()
	retryableErrors := []string{
		"connection reset",
		"timeout",
		"temporary failure",
		"network is unreachable",
		"no route to host",
		"connection refused",
		"i/o timeout",
	}

	for _, retryable := range retryableErrors {
		if strings.Contains(errStr, retryable) {
			return true
		}
	}
	return false
}
