package constants

const (
	API_GENERAL_BASE_URL = "https://api.torbox.app/v1"
	API_SEARCH_BASE_URL  = "https://search-api.torbox.app"
)

// GENERAL API
const (
	PATH_TORRENTS_GET_ACTIVE       = "api/torrents/mylist"
	PATH_TORRENTS_GET_DOWNLOAD_URL = "api/torrents/requestdl"
	PATH_TORRENTS_CREATE           = "api/torrents/createtorrent"
	PATH_TORRENTS_CONTROL_ACTIVE   = "api/torrents/controltorrent"

	PATH_TORRENTS_GET_QUEUED     = "api/queued/getqueued"
	PATH_TORRENTS_CONTROL_QUEUED = "api/queued/controlqueued"
)

// SEARCH API
const (
	PATH_SEARCH_TORRENTS = "torrents"
	PATH_SEARCH_META     = "meta"
)

type SeedSetting int

const (
	Auto SeedSetting = iota + 1
	Seed
	NoSeed
)
