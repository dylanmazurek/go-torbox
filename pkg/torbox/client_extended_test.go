package torbox_test

import (
	"fmt"
	"testing"

	"github.com/dylanmazurek/go-torbox/pkg/torbox"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/testutils"
	"github.com/stretchr/testify/assert"
)

func TestClient_CheckCachedTorrents(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	resp, err := client.General.CheckCachedTorrents([]string{"abc123hash"})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)
}

func TestClient_ExportTorrentData(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	data, err := client.General.ExportTorrentData(101, "json")
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "exported data content", *data)
}

func TestClient_MagnetToFile(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	data, err := client.General.MagnetToFile("magnet:?xt=urn:btih:example")
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, "base64encodedtorrentfile", *data)
}

func TestClient_GetUsenetList(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	downloads, err := client.General.GetUsenetList(nil)
	assert.NoError(t, err)
	assert.NotNil(t, downloads)
	assert.Len(t, downloads, 1)
	assert.Equal(t, int64(301), downloads[0].ID)
	assert.Equal(t, "Usenet Download", downloads[0].Name)
}

func TestClient_FetchUsenetDownloadUrl(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	url, err := client.General.FetchUsenetDownloadUrl(301, 1)
	assert.NoError(t, err)
	assert.NotNil(t, url)
	assert.Equal(t, "https://cdn.torbox.test/usenet/301/1", *url)
}

func TestClient_ControlUsenetDownload(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	err = client.General.ControlUsenetDownload(301, "pause")
	assert.NoError(t, err)
}

func TestClient_GetWebDownloadList(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	downloads, err := client.General.GetWebDownloadList(nil)
	assert.NoError(t, err)
	assert.NotNil(t, downloads)
	assert.Len(t, downloads, 1)
	assert.Equal(t, int64(401), downloads[0].ID)
	assert.Equal(t, "Web Download", downloads[0].Name)
}

func TestClient_FetchWebDownloadUrl(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	url, err := client.General.FetchWebDownloadUrl(401, 1)
	assert.NoError(t, err)
	assert.NotNil(t, url)
	assert.Equal(t, "https://cdn.torbox.test/web/401/1", *url)
}

func TestClient_ControlWebDownload(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	err = client.General.ControlWebDownload(401, "pause")
	assert.NoError(t, err)
}

func TestClient_GetStats(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	stats, err := client.General.GetStats()
	assert.NoError(t, err)
	assert.NotNil(t, stats)
	assert.Equal(t, int64(1024000000), stats.TotalDownloaded)
	assert.Equal(t, int64(512000000), stats.TotalUploaded)
}

func TestClient_GetNotifications(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	notifications, err := client.General.GetNotifications()
	assert.NoError(t, err)
	assert.NotNil(t, notifications)
	assert.Len(t, notifications, 1)
	assert.Equal(t, int64(1), notifications[0].ID)
	assert.Equal(t, "Download Complete", notifications[0].Title)
}

func TestClient_ClearNotifications(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	err = client.General.ClearNotifications(nil)
	assert.NoError(t, err)
}

func TestClient_GetRSSFeeds(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	feeds, err := client.General.GetRSSFeeds()
	assert.NoError(t, err)
	assert.NotNil(t, feeds)
	assert.Len(t, feeds, 1)
	assert.Equal(t, int64(1), feeds[0].ID)
	assert.Equal(t, "Example Feed", feeds[0].Name)
	assert.True(t, feeds[0].Active)
}

func TestClient_GetRSSFeedItems(t *testing.T) {
	mock := testutils.NewMockTorBox()
	defer mock.Close()

	client, err := torbox.New(
		t.Context(),
		torbox.WithAPIKey("test-api-key"),
		torbox.WithBaseURL(fmt.Sprintf("%s/v1", mock.URL())),
	)

	assert.NoError(t, err)

	items, err := client.General.GetRSSFeedItems(nil)
	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Len(t, items, 1)
	assert.Equal(t, int64(1), items[0].ID)
	assert.Equal(t, "Example Item", items[0].Title)
	assert.False(t, items[0].Downloaded)
}
