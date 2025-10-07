package models

import "github.com/dylanmazurek/go-torbox/pkg/torbox/constants"

type ControlActiveTorrentRequest struct {
	TorrentID int64                            `json:"torrent_id"`
	Operation constants.ControlActiveOperation `json:"operation"`
}

type ControlActiveTorrentResponse struct {
	BaseResponse
}
