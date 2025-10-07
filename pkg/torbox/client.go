package torbox

import (
	"context"
	"net/http"
	"time"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/general"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/search"
	"github.com/rs/zerolog/log"
)

type Client struct {
	General *general.GeneralService
	Search  *search.SearchService
}

func New(ctx context.Context, opts ...Option) (*Client, error) {
	log := log.Ctx(ctx).With().Str("component", "torbox").Logger()

	log.Info().Msg("initializing")

	clientOptions := defaultOptions()
	for _, opt := range opts {
		opt(&clientOptions)
	}

	httpAuthClient := &http.Client{
		Timeout: 60 * time.Second, // Increased timeout for file operations
		Transport: &addAuthHeaderTransport{
			T: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 5,
				IdleConnTimeout:     30 * time.Second,
				DisableKeepAlives:   false,
			},
			APIKey: clientOptions.apiKey,
		},
	}

	client := Client{
		General: general.New(*httpAuthClient, clientOptions.apiKey),
		Search:  search.New(*httpAuthClient),
	}

	return &client, nil
}
