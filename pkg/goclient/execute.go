package goclient

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
	"github.com/perimeterx/marshmallow"
	"github.com/rs/zerolog/log"
)

func (c *Client) Do(req *request.Request, obj any) error {
	return c.DoWithRetry(req, obj, 3)
}

func (c *Client) DoWithRetry(req *request.Request, obj any, maxRetries int) error {
	var lastErr error

	httpReq, err := req.ToHTTPRequest()
	if err != nil {
		return err
	}

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			delay := time.Duration(1<<(attempt-1)) * time.Second
			log.Debug().
				Int("attempt", attempt).
				Dur("delay", delay).
				Msg("retrying torbox API request after delay")

			time.Sleep(delay)
		}

		httpReq.Header.Add("Accept", "application/json")

		httpResponse, err := c.internalClient.Do(httpReq)
		if err != nil {
			lastErr = err
			log.Warn().
				Err(err).
				Int("attempt", attempt).
				Msg("torbox API request failed")

			if attempt < maxRetries && isRetryableNetworkError(err) {
				continue
			}

			log.Error().Err(err).Msg("failed to execute torbox request after retries")
			return err
		}

		if httpResponse != nil && httpResponse.StatusCode >= 400 {
			bodyBytes, _ := io.ReadAll(httpResponse.Body)
			httpResponse.Body.Close()

			var errResp models.BaseResponse
			unknownFields, err := marshmallow.Unmarshal(bodyBytes, &errResp, marshmallow.WithExcludeKnownFieldsFromMap(true))
			if err != nil {
				lastErr = err
				if attempt < maxRetries {
					log.Warn().
						Err(err).
						Int("attempt", attempt).
						Msg("failed to unmarshal torbox error response, retrying")

					continue
				}

				return err
			}

			if len(unknownFields) > 0 {
				log.Warn().Fields(unknownFields).Msgf("unknown fields in torbox error response")
			}

			log.Debug().
				Str("status", httpResponse.Status).
				Str("error", errResp.Error).
				Str("detail", errResp.Detail).
				Int("attempt", attempt).
				Msg("torbox API response error")

			// Handle rate limiting (429) with Retry-After header
			if httpResponse.StatusCode == 429 {
				retryAfter := httpResponse.Header.Get("Retry-After")
				if retryAfter != "" && attempt < maxRetries {
					if seconds, parseErr := strconv.Atoi(retryAfter); parseErr == nil {
						retryDelay := time.Duration(seconds) * time.Second
						log.Warn().
							Dur("retry_after", retryDelay).
							Int("attempt", attempt).
							Msg("rate limited by torbox API, waiting before retry")

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
						Msg("rate limited by torbox API, using exponential backoff")

					time.Sleep(delay)
					continue
				}
			}

			// Handle 5xx server errors with retry
			if httpResponse.StatusCode >= 500 && httpResponse.StatusCode < 600 && attempt < maxRetries {
				log.Warn().
					Int("status_code", httpResponse.StatusCode).
					Int("attempt", attempt).
					Msg("server error from torbox API, retrying")

				continue
			}

			// Create detailed error message from API response
			if errResp.Error != "" && errResp.Detail != "" {
				lastErr = fmt.Errorf("torbox API error: %s - %s (status: %d)", errResp.Error, errResp.Detail, httpResponse.StatusCode)
			} else if errResp.Error != "" {
				lastErr = fmt.Errorf("torbox API error: %s (status: %d)", errResp.Error, httpResponse.StatusCode)
			} else if errResp.Detail != "" {
				lastErr = fmt.Errorf("torbox API error: %s (status: %d)", errResp.Detail, httpResponse.StatusCode)
			} else {
				lastErr = fmt.Errorf("torbox server error (status: %d)", httpResponse.StatusCode)
			}
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
						Msg("failed to read torbox response body, retrying")

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
						Msg("failed to unmarshal torbox response, retrying")

					continue
				}
				return err
			}

			if len(unknownFields) > 0 {
				log.Warn().Fields(unknownFields).
					Msgf("unknown fields in torbox response")
			}
		}

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
