package general

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/aws/smithy-go/ptr"
	"github.com/dylanmazurek/go-torbox/pkg/goclient/request"
	"github.com/dylanmazurek/go-torbox/pkg/magnet"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

func (s *GeneralService) FetchDownloadUrl(torrentId int64, fileId int64) (*string, error) {
	var params url.Values
	params.Add("torrent_id", fmt.Sprintf("%d", torrentId))
	params.Add("file_id", fmt.Sprintf("%d", fileId))

	req, err := s.internalClient.NewRequest(http.MethodGet, constants.PATH_TORRENTS_GET_DOWNLOAD_URL, request.WithURLParams(&params))
	if err != nil {
		return nil, err
	}

	var resp struct {
		models.BaseResponse
		DownloadUrl string `json:"data"`
	}

	err = s.internalClient.Do(req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.DownloadUrl, nil
}

type CreateTorrentRequest struct {
	Magnet   *magnet.Magnet         `form:"magnet,omitempty"`
	File     []byte                 `form:"file,omitempty"`
	Seed     *constants.SeedSetting `form:"seed,omitempty"`
	AllowZip *bool                  `form:"allow_zip,omitempty"`
	Name     *string                `form:"name,omitempty"`
	AsQueued *bool                  `form:"as_queued,omitempty"`
}

func (s *GeneralService) CreateTorrent(r CreateTorrentRequest) (*models.Torrent, error) {
	var params = &url.Values{}
	var reqBody *bytes.Buffer

	reqOptions := []request.Option{}
	if r.File != nil {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "torrent.torrent")

		_, err := part.Write(r.File)
		if err != nil {
			return nil, err
		}

		writer.Close()
		params.Set("Content-Type", writer.FormDataContentType())

		reqBody = body
		reqOptions = append(reqOptions, request.WithFileBody(reqBody))
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
		reqOptions = append(reqOptions, request.WithFormBody(reqBody))
	}

	req, err := s.internalClient.NewRequest(http.MethodPost, constants.PATH_TORRENTS_CREATE, reqOptions...)
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
		return nil, err
	}

	return resp.Data, nil
}

func (s *GeneralService) ControlTorrent(id int64, operation string) error {
	activeTorrents, err := s.ListActiveTorrents(true)
	if err != nil {
		return err
	}

	isActive := slices.ContainsFunc(activeTorrents, func(t models.Torrent) bool {
		return t.ID == id
	})

	if isActive {
		var activeOperation constants.ControlActiveTorrentOperation
		err := activeOperation.FromString(operation)
		if err != nil {
			return err
		}

		return s.ControlActiveTorrent(id, activeOperation)
	}

	queuedTorrents, err := s.GetQueuedTorrents(ptr.Bool(true))
	if err != nil {
		return err
	}

	isQueued := slices.ContainsFunc(queuedTorrents, func(t models.QueuedDownload) bool {
		return t.ID == id
	})

	if isQueued {
		var controlQueuedRequest ControlQueuedTorrentRequest
		controlQueuedRequest.QueuedId = id
		err := controlQueuedRequest.Operation.FromString(operation)
		if err != nil {
			return err
		}

		return s.ControlQueuedTorrent(controlQueuedRequest)
	}

	return fmt.Errorf("torrent with ID %d is neither active nor queued", id)
}
