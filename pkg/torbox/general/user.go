package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetUser() (*models.User, error) {
	req, err := s.newRequest(http.MethodGet, constants.PATH_USER_ME, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetUserResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get user: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) RefreshToken() (*string, error) {
	req, err := s.newRequest(http.MethodPost, constants.PATH_USER_REFRESH_TOKEN, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.RefreshTokenResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to refresh token: %s", resp.Detail)
	}

	return &resp.Data.Token, nil
}

func (s *GeneralService) AddReferral(referralCode string) error {
	r := models.AddReferralRequest{
		ReferralCode: referralCode,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_USER_ADD_REFERRAL, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to add referral: %s", resp.Detail)
	}

	return nil
}
