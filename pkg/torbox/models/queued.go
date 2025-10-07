package models

import "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"

type GetQueuedTorrentsResponse struct {
	BaseResponse
	Data []QueuedDownload `json:"data"`
}

type QueuedDownload struct {
	ID          int64   `json:"id"`
	CreatedAt   string  `json:"created_at"`
	Magnet      string  `json:"magnet"`
	TorrentFile *string `json:"torrent_file"`
	Hash        string  `json:"hash"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
}

type ControlQueuedTorrentRequest struct {
	QueuedId  int64                            `json:"queued_id"`
	Operation constants.ControlQueuedOperation `json:"operation"`
}

type ControlQueuedTorrentResponse struct {
	BaseResponse
}
