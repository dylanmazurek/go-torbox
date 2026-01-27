package general

import (
	"context"
	"net/http"

	"github.com/dylanmazurek/go-torbox/pkg/goclient"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
)

type GeneralService struct {
	token string

	internalClient *goclient.Client
}

func New(ctx context.Context, internalClient http.Client, token string, baseUrl string) *GeneralService {
	url := constants.API_GENERAL_BASE_URL
	if baseUrl != "" {
		url = baseUrl
	}

	goClientOpts := []goclient.Option{
		goclient.WithBaseURL(url),
		goclient.WithHTTPClient(&internalClient),
	}

	newGoClient := goclient.NewClient(ctx, goClientOpts...)

	newGeneralService := &GeneralService{
		token: token,

		internalClient: newGoClient,
	}

	return newGeneralService
}
