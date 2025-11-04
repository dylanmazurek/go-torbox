package general

import (
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) GetActiveTorrents() ([]models.Torrent, error) {
	var params = &url.Values{}
	params.Add("bypass_cache", "true")

	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_TORRENTS_GET_ACTIVE, request.WithURLParams(params))
	if err != nil {
		return nil, err
	}

	var resp models.GetActiveTorrentsResponse
	err = s.internalClient.Do(req, &resp)
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

	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_TORRENTS_CONTROL_ACTIVE, request.WithJSONBody(r))
	if err != nil {
		return err
	}

	var resp models.ControlActiveTorrentResponse
	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return err
	}

	return nil
}
