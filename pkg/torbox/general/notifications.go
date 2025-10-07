package general

import (
	"fmt"
	"net/http"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetRSSNotifications() ([]models.Notification, error) {
	req, err := s.newRequest(http.MethodGet, constants.PATH_NOTIFICATIONS_RSS, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetNotificationsResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get RSS notifications: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) GetNotifications() ([]models.Notification, error) {
	req, err := s.newRequest(http.MethodGet, constants.PATH_NOTIFICATIONS_LIST, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetNotificationsResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get notifications: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) ClearNotifications() error {
	req, err := s.newRequest(http.MethodPost, constants.PATH_NOTIFICATIONS_CLEAR, nil, nil)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to clear notifications: %s", resp.Detail)
	}

	return nil
}
