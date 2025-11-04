package goclient

import (
	"context"
	"net/http"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
)

type Client struct {
	baseUrl string

	internalClient *http.Client
}

func NewClient(ctx context.Context, opts ...Option) *Client {
	clientOptions := defaultOptions()
	for _, opt := range opts {
		opt(&clientOptions)
	}

	newClient := &Client{
		baseUrl:        clientOptions.baseUrl,
		internalClient: clientOptions.internalClient,
	}

	return newClient
}

func (c *Client) NewRequest(m string, p string, opts ...request.Option) (*request.Request, error) {
	allOpts := []request.Option{
		request.WithBaseURL(c.baseUrl),
	}

	allOpts = append(allOpts, opts...)

	newRequest, err := request.NewRequest(m, p, allOpts...)

	return newRequest, err
}
