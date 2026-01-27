package general

import (
	"fmt"
	"net/http"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetStats() (*models.Stats, error) {
	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_STATS)
	if err != nil {
		return nil, err
	}

	var resp models.StatsResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get stats: %s", resp.Detail)
	}

	return resp.Data, nil
}
