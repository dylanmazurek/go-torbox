package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetQueuedTorrents(bypassCache bool) ([]models.QueuedDownload, error) {
	var params = &url.Values{}
	params.Add("bypass_cache", fmt.Sprintf("%t", bypassCache))

	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_TORRENTS_GET_QUEUED, request.WithURLParams(params))
	if err != nil {
		return nil, err
	}

	var resp models.GetQueuedTorrentsResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *GeneralService) ControlQueuedTorrent(queuedId int64, operation constants.ControlQueuedOperation) error {
	r := models.ControlQueuedTorrentRequest{
		QueuedId:  queuedId,
		Operation: operation,
	}

	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_TORRENTS_CONTROL_QUEUED, request.WithJSONBody(r))
	if err != nil {
		return err
	}

	var resp models.ControlQueuedTorrentResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return err
	}

	return nil
}
