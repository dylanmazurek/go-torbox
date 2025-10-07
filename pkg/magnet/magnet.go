package magnet

import (
	"errors"
	"net/url"
	"strings"
)

type Magnet struct {
	Hash        string   `json:"hash"`
	DisplayName string   `json:"display_name"`
	Trackers    []string `json:"trackers"`

	originalUrl string `json:"-"`
}

func NewMagnet(magnetUrl string) (*Magnet, error) {
	parsedMagnetUrl, err := url.Parse(magnetUrl)
	if err != nil {
		return nil, err
	}

	if parsedMagnetUrl == nil || parsedMagnetUrl.Scheme != "magnet" {
		return nil, errors.New("invalid magnet link: not a magnet scheme")
	}

	queryParams := parsedMagnetUrl.Query()

	exactTopic := queryParams.Get("xt")
	var magnetHash string
	if exactTopic == "" {
		return nil, errors.New("invalid magnet link: missing or malformed xt parameter")
	}

	if strings.HasPrefix(exactTopic, BitTorrentInfoHashPrefix) {
		magnetHash = strings.TrimPrefix(exactTopic, BitTorrentInfoHashPrefix)
	}

	displayName := queryParams.Get("dn")
	if displayName == "" {
		displayName = "unknown"
	}

	newMagnet := &Magnet{
		Hash:        magnetHash,
		DisplayName: displayName,

		originalUrl: magnetUrl,
	}

	return newMagnet, nil
}

func (m *Magnet) GetUrl() *string {
	if m.originalUrl != "" {
		return &m.originalUrl
	}

	return nil
}

const (
	BitTorrentInfoHashPrefix = "urn:btih:"
)
