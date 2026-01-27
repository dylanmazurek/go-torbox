package general

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

// GetActiveTorrent retrieves information about a specific active torrent by its hash.
// h is the hash of the torrent to retrieve.
func (s *GeneralService) GetActiveTorrent(h string) (*models.Torrent, error) {
	var params url.Values = make(url.Values)
	params.Add("hash", h)

	req, err := s.internalClient.NewRequest(
		http.MethodGet,
		constants.PATH_TORRENTS_INFO,
		request.WithURLParams(&params),
	)
	if err != nil {
		return nil, err
	}

	var resp struct {
		models.BaseResponse
		Data *models.Torrent `json:"data"`
	}

	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get torrent info: %s", resp.Detail)
	}

	return resp.Data, nil
}

// ListActiveTorrents retrieves a list of active torrents.
// b is a boolean indicating whether to bypass the cache.
func (s *GeneralService) ListActiveTorrents(b bool) ([]models.Torrent, error) {
	var params url.Values = make(url.Values)
	if b {
		params.Add("bypass_cache", "true")
	}

	req, err := s.internalClient.NewRequest(
		http.MethodGet,
		constants.PATH_TORRENTS_GET_ACTIVE,
		request.WithURLParams(&params),
	)
	if err != nil {
		return nil, err
	}

	var resp struct {
		models.BaseResponse
		Data []models.Torrent `json:"data"`
	}

	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (s *GeneralService) ControlActiveTorrent(torrentId int64, op constants.ControlActiveTorrentOperation) error {
	r := models.ControlActiveTorrentRequest{
		TorrentID: torrentId,
		Operation: op,
	}

	req, err := s.internalClient.NewRequest(
		http.MethodPost,
		constants.PATH_TORRENTS_CONTROL_ACTIVE,
		request.WithJSONBody(r),
	)
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
