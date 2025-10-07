package general

import (
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetActiveTorrents() ([]models.Torrent, error) {
	var params = &url.Values{}
	params.Add("bypass_cache", "true")

	req, err := s.newRequest(http.MethodGet, constants.PATH_TORRENTS_GET_ACTIVE, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetActiveTorrentsResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *GeneralService) ControlActiveTorrent(torrentId int64, operation constants.ControlActiveOperation) error {
	r := models.ControlActiveTorrentRequest{
		TorrentID: torrentId,
		Operation: operation,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_TORRENTS_CONTROL_ACTIVE, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.ControlActiveTorrentResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	return nil
}
