package goclient

import "net/http"

type options struct {
	internalClient *http.Client
	baseUrl        string
}

func defaultOptions() options {
	defaultOptions := options{
		internalClient: http.DefaultClient,
	}

	return defaultOptions
}

type Option func(*options)

func WithHTTPClient(c *http.Client) Option {
	return func(o *options) {
		o.internalClient = c
	}
}

func WithBaseURL(u string) Option {
	return func(o *options) {
		o.baseUrl = u
	}
}
