package general

import (
	"fmt"
	"net/http"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetUser() (*models.User, error) {
	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_USER_ME)
	if err != nil {
		return nil, err
	}

	var resp models.GetUserResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get user: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) RefreshToken() (*string, error) {
	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_USER_REFRESH_TOKEN)
	if err != nil {
		return nil, err
	}

	var resp models.RefreshTokenResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to refresh token: %s", resp.Detail)
	}

	return &resp.Data.Token, nil
}
