package general

import (
	"fmt"
	"net/http"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetNotifications() ([]models.Notification, error) {
	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_NOTIFICATIONS_LIST)
	if err != nil {
		return nil, err
	}

	var resp models.NotificationsResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get notifications: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) ClearNotifications(notificationId *int64) error {
	r := models.ClearNotificationRequest{
		ID: notificationId,
	}

	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_NOTIFICATIONS_CLEAR, request.WithJSONBody(r))
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to clear notifications: %s", resp.Detail)
	}

	return nil
}
