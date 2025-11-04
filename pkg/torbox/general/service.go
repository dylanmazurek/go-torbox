package general

import (
	"net/http"

	"github.com/dylanmazurek/go-torbox/pkg/goclient"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
)

type GeneralService struct {
	token string

	internalClient *goclient.Client
}

func New(internalClient http.Client, token string) *GeneralService {
	goClientOpts := []goclient.Option{
		goclient.WithBaseURL(constants.API_GENERAL_BASE_URL),
		goclient.WithHTTPClient(&internalClient),
	}

	newGoClient := goclient.NewClient(nil, goClientOpts...)

	newGeneralService := &GeneralService{
		token: token,

		internalClient: newGoClient,
	}

	return newGeneralService
}
