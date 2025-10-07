package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) AddRSS(r models.AddRSSRequest) (*models.RSSFeed, error) {
	req, err := s.newRequest(http.MethodPost, constants.PATH_RSS_ADD, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return nil, err
	}

	var resp models.AddRSSResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to add RSS: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) ControlRSS(rssId int64, operation constants.ControlRSSOperation) error {
	r := models.ControlRSSRequest{
		RSSID:     rssId,
		Operation: operation,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_RSS_CONTROL, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to control RSS: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) ModifyRSS(r models.ModifyRSSRequest) (*models.RSSFeed, error) {
	req, err := s.newRequest(http.MethodPost, constants.PATH_RSS_MODIFY, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return nil, err
	}

	var resp models.ModifyRSSResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to modify RSS: %s", resp.Detail)
	}

	return resp.Data, nil
}
