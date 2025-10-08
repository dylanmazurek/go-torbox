package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) AuthorizeGoogleDrive(code string) error {
	r := models.IntegrationAuthRequest{
		Code: code,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_INTEGRATION_GOOGLEDRIVE, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to authorize Google Drive: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) AuthorizeDropbox(code string) error {
	r := models.IntegrationAuthRequest{
		Code: code,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_INTEGRATION_DROPBOX, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to authorize Dropbox: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) AuthorizeOneDrive(code string) error {
	r := models.IntegrationAuthRequest{
		Code: code,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_INTEGRATION_ONEDRIVE, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to authorize OneDrive: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) AuthorizeGofile(apiKey string) error {
	r := models.IntegrationAuthRequest{
		APIKey: apiKey,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_INTEGRATION_GOFILE, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to authorize Gofile: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) Authorize1Fichier(apiKey string) error {
	r := models.IntegrationAuthRequest{
		APIKey: apiKey,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_INTEGRATION_1FICHIER, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to authorize 1Fichier: %s", resp.Detail)
	}

	return nil
}

func (s *GeneralService) GetIntegrationJobs() ([]models.IntegrationJob, error) {
	req, err := s.newRequest(http.MethodGet, constants.PATH_INTEGRATION_JOBS, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetIntegrationJobsResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get integration jobs: %s", resp.Detail)
	}

	return resp.Data, nil
}
