package general

import (
	"fmt"
	"net/http"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetStats() (*models.Stats, error) {
	req, err := s.newRequest(http.MethodGet, constants.PATH_STATS, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetStatsResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get stats: %s", resp.Detail)
	}

	return resp.Data, nil
}
