package models

import (
	"github.com/dylanmazurek/go-torbox/pkg/magnet"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
)

type GetActiveTorrentsResponse struct {
	BaseResponse
	Data []Torrent `json:"data"`
}

type GetDownloadUrlResponse struct {
	BaseResponse
	DownloadUrl string `json:"data"`
}

type CreateTorrentRequest struct {
	Magnet   *magnet.Magnet         `form:"magnet,omitempty"`
	File     []byte                 `form:"file,omitempty"`
	Seed     *constants.SeedSetting `form:"seed,omitempty"`
	AllowZip *bool                  `form:"allow_zip,omitempty"`
	Name     *string                `form:"name,omitempty"`
	AsQueued *bool                  `form:"as_queued,omitempty"`
}

type CreateTorrentResponse struct {
	BaseResponse
	Data *Torrent `json:"data"`
}

type GetMetaResponse struct {
	BaseResponse
	Data Torrent `json:"data"`
}

type GetTorrentResponse struct {
	BaseResponse
	Data struct {
		Metadata Metadata  `json:"metadata"`
		Torrents []Torrent `json:"torrents"`
	} `json:"data"`
}

type CheckCachedResponse struct {
	BaseResponse
	Data *CacheCheckResponse `json:"data"`
}

type CacheCheckResponse struct {
	Name   string `json:"name"`
	Size   int64  `json:"size"`
	Hash   string `json:"hash"`
	Cached bool   `json:"cached"`
}

type TorrentInfoResponse struct {
	BaseResponse
	Data *Torrent `json:"data"`
}

type ExportDataResponse struct {
	BaseResponse
	Data string `json:"data"`
}

type SearchTorrentsResponse struct {
	BaseResponse
	Data []Torrent `json:"data"`
}

type StoreSearchRequest struct {
	Query string `json:"query"`
}

