package constants

type TorrentState string

// processing states
const (
	// The torrent is checking resumable data.
	TorrentStateCheckingResumeData TorrentState = "checkingResumeData"
	// The torrent is checking files.
	TorrentStateChecking = "checking"
	// The torrent is downloading metadata from the hoard.
	TorrentStateMetaDL = "metaDL"
	// The torrent is paused.
	TorrentStatePaused = "paused"
)

// downloading states
const (
	// The torrent is currently downloading.
	TorrentStateDownloading = "downloading"
	// The torrent is trying to download, but there are no seeds connected to download from.
	TorrentStateStalledNoSeeds = "stalled (no seeds)"
	// The torrent is trying to download, but there are no peers connected to download from.
	TorrentStateStalledDL = "stalledDL"
)

// completed states
const (
	// The torrent is completely downloaded. Do not use this for download completion status.
	TorrentStateCompleted = "completed"
	// The torrent is cached from the server.
	TorrentStateCached = "cached"
)

// uploading states
const (
	// The torrent is currently seeding.
	TorrentStateUploading = "uploading"
	// The torrent is currently seeding, but there are no peers connected to upload to.
	TorrentStateUploadingNoPeers = "uploading (no peers)"
)

const (
	// The torrent state is unknown.
	TorrentStateUnknown = "unknown"
)

func (t TorrentState) IsComplete() bool {
	switch t {
	case TorrentStateCompleted,
		TorrentStateCached,
		TorrentStateUploading,
		TorrentStateUploadingNoPeers:

		return true
	default:
		return false
	}
}
