package testutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"

	"github.com/dylanmazurek/go-torbox/pkg/torbox/constants"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
)

// MockTorBox represents a mock TorBox API server
type MockTorBox struct {
	Server *httptest.Server
	Mux    *http.ServeMux

	ActiveTorrents []models.Torrent
	QueuedTorrents []models.QueuedDownload

	ActiveControlRequests []models.ControlActiveTorrentRequest
	QueuedControlRequests []QueuedControlRequest
}

type QueuedControlRequest struct {
	QueuedID  int64                                   `json:"queued_id"`
	Operation constants.ControlQueuedTorrentOperation `json:"operation"`
}

// NewMockTorBox creates a new mock TorBox API server
func NewMockTorBox() *MockTorBox {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mock := &MockTorBox{
		Server:         server,
		Mux:            mux,
		ActiveTorrents: []models.Torrent{defaultActiveTorrent()},
		QueuedTorrents: []models.QueuedDownload{defaultQueuedDownload()},
	}

	mock.setupRoutes()
	return mock
}

// Close shuts down the mock server
func (m *MockTorBox) Close() {
	m.Server.Close()
}

// URL returns the base URL of the mock server
func (m *MockTorBox) URL() string {
	return m.Server.URL
}

func (m *MockTorBox) setupRoutes() {
	m.Mux.HandleFunc("/v1/api/user/me", m.handleUserMe)
	m.Mux.HandleFunc("/v1/api/torrents/mylist", m.handleTorrentsList)
	m.Mux.HandleFunc("/v1/api/torrents/torrentinfo", m.handleTorrentInfo)
	m.Mux.HandleFunc("/v1/api/torrents/requestdl", m.handleTorrentDownloadURL)
	m.Mux.HandleFunc("/v1/api/torrents/createtorrent", m.handleTorrentCreate)
	m.Mux.HandleFunc("/v1/api/torrents/controltorrent", m.handleControlActiveTorrent)
	m.Mux.HandleFunc("/v1/api/queued/getqueued", m.handleQueuedList)
	m.Mux.HandleFunc("/v1/api/queued/controlqueued", m.handleControlQueuedTorrent)
}

func (m *MockTorBox) handleUserMe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authID := "auth0|123456"
	baseEmail := "test@example.com"
	vendorID := int64(1)

	user := models.User{
		ID:            123,
		Email:         "test@example.com",
		Plan:          constants.PlanTypePro,
		PremiumExpiry: time.Now().Add(30 * 24 * time.Hour).Format("2006-01-02"),
		CooldownUntil: time.Now().Format("2006-01-02 15:04:05"),
		Customer:      "cus_123456",
		Server:        1,
		IsSubscribed:  true,
		UserReferral:  "REF123",
		Auth0ID:       &authID,
		BaseEmail:     &baseEmail,
		IsVendor:      true,
		VendorID:      &vendorID,
	}

	response := models.GetUserResponse{
		BaseResponse: models.BaseResponse{
			Success: true,
			Detail:  "User details retrieved successfully",
		},
		Data: &user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (m *MockTorBox) handleTorrentsList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]any{
		"success": true,
		"detail":  "Torrents retrieved",
		"data":    m.ActiveTorrents,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (m *MockTorBox) handleTorrentInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	hash := r.URL.Query().Get("hash")
	for _, t := range m.ActiveTorrents {
		if t.Hash != hash {
			continue
		}

		response := map[string]any{
			"success": true,
			"detail":  "Torrent details retrieved",
			"data": map[string]any{
				"id":           t.ID,
				"hash":         t.Hash,
				"server":       t.Server,
				"auth_id":      t.AuthID,
				"name":         t.Name,
				"magnet":       t.Magnet,
				"size":         t.Size,
				"eta":          t.ETA,
				"torrent_file": t.TorrentFile,
				"cached":       t.Cached,
				"owner":        t.Owner,
				"allow_zipped": t.AllowZipped,
				"short_name":   t.ShortName,
				"created_at":   formatTime(t.CreatedAt),
				"updated_at":   formatTime(t.UpdatedAt),
				"expires_at":   formatTime(t.ExpiresAt),
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	http.Error(w, "torrent not found", http.StatusNotFound)
}

func (m *MockTorBox) handleTorrentDownloadURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	torrentID := r.URL.Query().Get("torrent_id")
	fileID := r.URL.Query().Get("file_id")

	downloadURL := "https://cdn.torbox.test/download/" + torrentID + "/" + fileID

	response := map[string]any{
		"success": true,
		"detail":  "Download URL generated",
		"data":    downloadURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (m *MockTorBox) handleTorrentCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var name string = "New Torrent"
	var magnet string

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		if err := r.ParseForm(); err == nil {
			name = r.FormValue("name")
			magnet = r.FormValue("magnet")
		}
	}

	newID := int64(len(m.ActiveTorrents) + 200)
	now := time.Now().UTC()
	magnetHash := "hash" + strconv.FormatInt(newID, 10)

	if magnet != "" {
		const hashPrefix = "urn:btih:"
		if idx := strings.Index(magnet, hashPrefix); idx >= 0 {
			hashStart := idx + len(hashPrefix)
			hashEnd := strings.Index(magnet[hashStart:], "&")
			if hashEnd == -1 {
				hashEnd = len(magnet)
			} else {
				hashEnd = hashStart + hashEnd
			}

			magnetHash = magnet[hashStart:hashEnd]
		}
	}

	torrent := models.Torrent{
		ID:          newID,
		Hash:        magnetHash,
		Server:      1,
		AuthID:      "auth0|123456",
		Name:        name,
		Magnet:      magnet,
		Size:        1024,
		ETA:         0,
		TorrentFile: magnet == "",
		Cached:      false,
		Owner:       "test_user",
		AllowZipped: true,
		ShortName:   name,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		ExpiresAt:   &now,
	}

	m.ActiveTorrents = append(m.ActiveTorrents, torrent)

	response := map[string]any{
		"success": true,
		"detail":  "Torrent created",
		"data":    torrent,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (m *MockTorBox) handleControlActiveTorrent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body models.ControlActiveTorrentRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err == nil {
		m.ActiveControlRequests = append(m.ActiveControlRequests, body)
	}

	response := map[string]any{
		"success": true,
		"detail":  "Active torrent control acknowledged",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (m *MockTorBox) handleQueuedList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]any{
		"success": true,
		"detail":  "Queued torrents retrieved",
		"data":    m.QueuedTorrents,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (m *MockTorBox) handleControlQueuedTorrent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body QueuedControlRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err == nil {
		m.QueuedControlRequests = append(m.QueuedControlRequests, body)
	}

	response := map[string]any{
		"success": true,
		"detail":  "Queued torrent control acknowledged",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func defaultActiveTorrent() models.Torrent {
	now := time.Now().UTC()
	return models.Torrent{
		ID:          101,
		Hash:        "abc123hash",
		Server:      1,
		AuthID:      "auth0|123456",
		Name:        "Ubuntu Linux ISO",
		Magnet:      "magnet:?xt=urn:btih:abc123hash",
		Size:        2048000000,
		ETA:         0,
		TorrentFile: false,
		Cached:      true,
		Owner:       "test_user",
		AllowZipped: true,
		ShortName:   "Ubuntu",
		CreatedAt:   &now,
		UpdatedAt:   &now,
		ExpiresAt:   &now,
	}
}

func defaultQueuedDownload() models.QueuedDownload {
	magnet := "magnet:?xt=urn:btih:def456hash"
	torrentFile := ""

	return models.QueuedDownload{
		ID:          202,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
		Magnet:      &magnet,
		TorrentFile: &torrentFile,
		Hash:        "def456hash",
		Name:        "Fedora ISO",
		Type:        "torrent",
	}
}

func formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}

	return t.UTC().Format("2006-01-02T15:04:05Z")
}
