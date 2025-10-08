package general

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) CreateTorrent(r models.CreateTorrentRequest) (*models.Torrent, error) {
	var params = &url.Values{}
	var reqBody *bytes.Buffer

	if r.File != nil {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "torrent.torrent")

		_, err := part.Write(r.File)
		if err != nil {
			return nil, err
		}

		writer.Close()
		params.Set("bodyType", "file")
		params.Set("Content-Type", writer.FormDataContentType())

		reqBody = body
	} else if r.Magnet != nil {
		body := &bytes.Buffer{}

		form := url.Values{}
		form.Set("magnet", *r.Magnet.GetUrl())
		form.Set("name", *r.Name)

		if r.AsQueued != nil {
			form.Set("as_queued", fmt.Sprintf("%t", *r.AsQueued))
		}

		_, err := body.WriteString(form.Encode())
		if err != nil {
			return nil, err
		}

		params.Set("bodyType", "form")
		params.Set("Content-Type", "application/x-www-form-urlencoded")

		reqBody = body
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_TORRENTS_CREATE, params, reqBody)
	if err != nil {
		return nil, err
	}

	var resp models.CreateTorrentResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, err
	}

	return resp.Data, nil
}

func (s *GeneralService) GetDownloadUrl(torrentId int64, fileId int64) (*string, error) {
	params := &url.Values{}
	params.Add("torrent_id", fmt.Sprintf("%d", torrentId))
	params.Add("file_id", fmt.Sprintf("%d", fileId))
	params.Add("token", "")

	req, err := s.newRequest(http.MethodGet, constants.PATH_TORRENTS_GET_DOWNLOAD_URL, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GetDownloadUrlResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		if resp.Error != "" && resp.Detail != "" {
			return nil, fmt.Errorf("failed to get download URL: %s - %s", resp.Error, resp.Detail)
		} else if resp.Error != "" {
			return nil, fmt.Errorf("failed to get download URL: %s", resp.Error)
		} else if resp.Detail != "" {
			return nil, fmt.Errorf("failed to get download URL: %s", resp.Detail)
		} else {
			return nil, fmt.Errorf("failed to get download URL: unknown error")
		}
	}

	return &resp.DownloadUrl, nil
}

func (s *GeneralService) ControlAnyTorrent(id int64, operation string) error {
	activeTorrents, err := s.GetActiveTorrents()
	if err != nil {
		return err
	}

	isActive := slices.ContainsFunc(activeTorrents, func(t models.Torrent) bool {
		return t.ID == id
	})

	if isActive {
		activeOperation := constants.ControlActiveOperation(operation)

		return s.ControlActiveTorrent(id, activeOperation)
	}

	queuedTorrents, err := s.GetQueuedTorrents()
	if err != nil {
		return err
	}

	isQueued := slices.ContainsFunc(queuedTorrents, func(t models.QueuedDownload) bool {
		return t.ID == id
	})

	if isQueued {
		queuedOperation := constants.ControlQueuedOperation(operation)

		return s.ControlQueuedTorrent(id, queuedOperation)
	}

	return fmt.Errorf("torrent with ID %d is neither active nor queued", id)
}

func (s *GeneralService) CheckCached(hash string) (*models.CacheCheckResponse, error) {
	params := &url.Values{}
	params.Add("hash", hash)
	params.Add("format", "list")

	req, err := s.newRequest(http.MethodGet, constants.PATH_TORRENTS_CHECK_CACHED, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.CheckCachedResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to check cache: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) GetTorrentInfo(hash string) (*models.Torrent, error) {
	params := &url.Values{}
	params.Add("hash", hash)

	req, err := s.newRequest(http.MethodGet, constants.PATH_TORRENTS_INFO, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.TorrentInfoResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to get torrent info: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) ExportData() (string, error) {
	req, err := s.newRequest(http.MethodGet, constants.PATH_TORRENTS_EXPORT_DATA, nil, nil)
	if err != nil {
		return "", err
	}

	var resp models.ExportDataResponse
	err = s.do(req, &resp)
	if err != nil {
		return "", err
	}

	if !resp.Success {
		return "", fmt.Errorf("failed to export data: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) SearchTorrents(query string) ([]models.Torrent, error) {
	params := &url.Values{}
	params.Add("query", query)

	req, err := s.newRequest(http.MethodGet, constants.PATH_TORRENTS_SEARCH, params, nil)
	if err != nil {
		return nil, err
	}

	var resp models.SearchTorrentsResponse
	err = s.do(req, &resp)
	if err != nil {
		return nil, err
	}

	if !resp.Success {
		return nil, fmt.Errorf("failed to search torrents: %s", resp.Detail)
	}

	return resp.Data, nil
}

func (s *GeneralService) StoreSearch(query string) error {
	r := models.StoreSearchRequest{
		Query: query,
	}

	req, err := s.newRequest(http.MethodPost, constants.PATH_TORRENTS_STORE_SEARCH, &url.Values{"bodyType": {"json"}}, r)
	if err != nil {
		return err
	}

	var resp models.BaseResponse
	err = s.do(req, &resp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("failed to store search: %s", resp.Detail)
	}

	return nil
}
