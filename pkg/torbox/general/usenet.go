package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) CreateUsenetDownload(r models.CreateUsenetRequest) (*models.UsenetDownload, error) {
	req, err := s.newRequest(http.MethodPost, constants.PATH_USENET_CREATE, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return nil, err
	}

	var resp models.CreateUsenetResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to create usenet download: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) GetUsenetList() ([]models.UsenetDownload, error) {
	params := &url.Values{}
	params.Add("bypass_cache", "true")

	req, err := s.newRequest(http.MethodGet, constants.PATH_USENET_GET_LIST, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetUsenetListResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get usenet list: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) ControlUsenetDownload(usenetId int64, operation constants.ControlUsenetOperation) error {
	r := models.ControlUsenetRequest{
		UsenetID:  usenetId,
		Operation: operation,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_USENET_CONTROL, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to control usenet download: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) GetUsenetDownloadUrl(usenetId int64, fileId int64) (*string, error) {
	params := &url.Values{}
	params.Add("usenet_id", fmt.Sprintf("%d", usenetId))
	params.Add("file_id", fmt.Sprintf("%d", fileId))
	params.Add("token", "")

	req, err := s.newRequest(http.MethodGet, constants.PATH_USENET_GET_DOWNLOAD, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetDownloadUrlResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		if resp.Error != "" && resp.Detail != "" {
			return nil, fmt.Errorf("failed to get download URL: %s - %s", resp.Error, resp.Detail)
		} else if resp.Error != "" {
			return nil, fmt.Errorf("failed to get download URL: %s", resp.Error)
		} else if resp.Detail != "" {
			return nil, fmt.Errorf("failed to get download URL: %s", resp.Detail)
		} else {
			return nil, fmt.Errorf("failed to get download URL: unknown error")
		}
	}

	return &resp.DownloadUrl, nil
}

func (s *GeneralService) CheckUsenetCached(hash string) (*models.CacheCheckResponse, error) {
	params := &url.Values{}
	params.Add("hash", hash)
	params.Add("format", "list")

	req, err := s.newRequest(http.MethodGet, constants.PATH_USENET_CHECK_CACHED, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.CheckCachedResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to check usenet cache: %s", resp.Detail)
	}

	return resp.Data, nil
}
