package general

import (
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetQueuedTorrents() ([]models.QueuedDownload, error) {
	var params = &url.Values{}
	params.Add("bypass_cache", "true")

	req, err := s.newRequest(http.MethodGet, constants.PATH_TORRENTS_GET_QUEUED, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetQueuedTorrentsResponse
	err = s.do(req, &resp)
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

	req, err := s.newRequest(http.MethodPost, constants.PATH_TORRENTS_CONTROL_QUEUED, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.ControlQueuedTorrentResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	return nil
}
