package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetQueuedTorrents(bypassCache *bool) ([]models.QueuedDownload, error) {
	var params url.Values = make(url.Values)
	if bypassCache != nil {
		params.Add("bypass_cache", fmt.Sprintf("%t", *bypassCache))
	}

	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_TORRENTS_GET_QUEUED, request.WithURLParams(&params))
	if err != nil {
		return nil, err
	}

	var resp struct {
		models.BaseResponse
		Data []models.QueuedDownload `json:"data"`
	}

	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

type ControlQueuedTorrentRequest struct {
	QueuedId  int64                                   `json:"queued_id"`
	Operation constants.ControlQueuedTorrentOperation `json:"operation"`
}

func (s *GeneralService) ControlQueuedTorrent(r ControlQueuedTorrentRequest) error {
	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_TORRENTS_CONTROL_QUEUED, request.WithJSONBody(r))
	if err != nil {
		return err
	}

	var resp struct {
		models.BaseResponse
	}

	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return err
	}

	return nil
}
