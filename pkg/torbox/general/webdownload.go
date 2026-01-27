package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) CreateWebDownload(r models.CreateWebDownloadRequest) (*models.WebDownload, error) {
	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_WEBDL_CREATE, request.WithJSONBody(r))
	if err != nil {
		return nil, err
	}

	var resp models.CreateWebDownloadResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to create web download: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) GetWebDownloadList(bypassCache *bool) ([]models.WebDownload, error) {
	params := make(url.Values)
	if bypassCache != nil {
		params.Add("bypass_cache", fmt.Sprintf("%t", *bypassCache))
	}

	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_WEBDL_GET_LIST, request.WithURLParams(&params))
	if err != nil {
		return nil, err
	}

	var resp models.GetWebDownloadListResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get web download list: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) FetchWebDownloadUrl(webId int64, fileId int64) (*string, error) {
	params := make(url.Values)
	params.Add("web_id", fmt.Sprintf("%d", webId))
	params.Add("file_id", fmt.Sprintf("%d", fileId))

	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_WEBDL_GET_DOWNLOAD, request.WithURLParams(&params))
	if err != nil {
		return nil, err
	}

	var resp struct {
		models.BaseResponse
		DownloadUrl string `json:"data"`
	}

	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.DownloadUrl, nil
}

func (s *GeneralService) ControlWebDownload(webId int64, operation string) error {
	r := models.ControlWebDownloadRequest{
		WebID:     webId,
		Operation: operation,
	}

	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_WEBDL_CONTROL, request.WithJSONBody(r))
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to control web download: %s", resp.Detail)
	}

	return nil
}
