package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) CreateWebDownload(r models.CreateWebDownloadRequest) (*models.WebDownload, error) {
	req, err := s.newRequest(http.MethodPost, constants.PATH_WEBDL_CREATE, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return nil, err
	}

	var resp models.CreateWebDownloadResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to create web download: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) ControlWebDownload(webId int64, operation constants.ControlWebDownloadOperation) error {
	r := models.ControlWebDownloadRequest{
		WebID:     webId,
		Operation: operation,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_WEBDL_CONTROL, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to control web download: %s", resp.Detail)
	}

	return nil
}
