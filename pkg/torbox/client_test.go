package torbox_test

import (
	"fmt"
	"testing"

	"github.com/dylanmazurek/go-torbox/pkg/magnet"
	"github.com/dylanmazurek/go-torbox/pkg/torbox"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/general"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/testutils"
	"github.com/stretchr/testify/assert"
)

func TestClient_UserMe(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	user, err := client.General.GetUser()
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, constants.PlanTypePro, user.Plan)
}

func TestClient_TorrentsList(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	torrents, err := client.General.ListActiveTorrents(false)
	assert.NoError(t, err)
	assert.NotNil(t, torrents)
	assert.Len(t, torrents, 1)
	assert.Equal(t, "Ubuntu Linux ISO", torrents[0].Name)
}

func TestClient_GetActiveTorrent(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	torrent, err := client.General.GetActiveTorrent("abc123hash")
	assert.NoError(t, err)
	assert.NotNil(t, torrent)
	assert.Equal(t, int64(101), torrent.ID)
	assert.Equal(t, "Ubuntu Linux ISO", torrent.Name)
	assert.NotNil(t, torrent.CreatedAt)
}

func TestClient_CreateTorrentWithMagnet(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	mag, err := magnet.NewMagnet("magnet:?xt=urn:btih:feedface&dn=Example")
	assert.NoError(t, err)

	name := "Example Torrent"
	torrent, err := client.General.CreateTorrent(general.CreateTorrentRequest{
		Magnet: mag,
		Name:   &name,
	})

	assert.NoError(t, err)
	assert.NotNil(t, torrent)
	assert.Equal(t, name, torrent.Name)
	assert.Equal(t, mag.Hash, torrent.Hash)
}

func TestClient_FetchDownloadUrl(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	url, err := client.General.FetchDownloadUrl(101, 1)
	assert.NoError(t, err)
	assert.NotNil(t, url)
	assert.Equal(t, "https://cdn.torbox.test/download/101/1", *url)
}

func TestClient_ControlTorrent_Active(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	err = client.General.ControlTorrent(101, "pause")
	assert.NoError(t, err)
	assert.Len(t, mock.ActiveControlRequests, 1)
	assert.Equal(t, int64(101), mock.ActiveControlRequests[0].TorrentID)
	assert.Equal(t, constants.ControlActiveOperationPause, mock.ActiveControlRequests[0].Operation)
}

func TestClient_ControlTorrent_Queued(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	err = client.General.ControlTorrent(202, "start")
	assert.NoError(t, err)
	assert.Len(t, mock.QueuedControlRequests, 1)
	assert.Equal(t, int64(202), mock.QueuedControlRequests[0].QueuedID)
	assert.Equal(t, constants.ControlQueuedOperationStart, mock.QueuedControlRequests[0].Operation)
}
