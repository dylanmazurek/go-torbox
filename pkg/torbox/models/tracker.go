package models

type TrackerDetails struct {
	InactiveCheck   int64  `json:"inactive_check"`
	LongTermSeeding bool   `json:"long_term_seeding"`
	TrackerMessage  string `json:"tracker_message"`
	SeedTorrent     bool   `json:"seed_torrent"`
	Active          bool   `json:"active"`

	Availability      float64 `json:"availability"`
	Ratio             float64 `json:"ratio"`
	Tracker           string  `json:"tracker"`
	Seeds             int64   `json:"seeds"`
	Peers             int64   `json:"peers"`
	LastKnownSeeders  int64   `json:"last_known_seeders"`
	LastKnownLeechers int64   `json:"last_known_leechers"`
}
