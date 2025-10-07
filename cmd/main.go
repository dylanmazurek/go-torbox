package main

import (
	"context"
	"os"

	"github.com/dylanmazurek/go-torbox/internal/logger"
	"github.com/dylanmazurek/go-torbox/pkg/torbox"
	"github.com/dylanmazurek/go-torbox/pkg/torbox/models"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	log.Logger = logger.New(ctx)

	apiKey, tokenExists := os.LookupEnv("TORBOX_API_KEY")
	if !tokenExists {
		log.Fatal().Msg("TORBOX_API_KEY environment variable is not set")
	}

	args := []torbox.Option{
		torbox.WithAPIKey(apiKey),
	}

	client, err := torbox.New(ctx, args...)
	if err != nil {
		panic(err)
	}

	activeTorrents, err := client.General.GetActiveTorrents()
	if err != nil {
		log.Error().Err(err).Msg("failed to get download URL")
		return
	}

	printTorrents(activeTorrents)
}

func printTorrents(torrents []models.Torrent) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Status", "Size", "Progress", "Download Speed", "Upload Speed", "Ratio"})

	for _, torrent := range torrents {
		t.AppendRow([]interface{}{
			torrent.ID,
			torrent.Name,
			torrent.DownloadState,
			torrent.Size,
			torrent.Progress,
			torrent.DownloadSpeed,
			torrent.UploadSpeed,
			torrent.Ratio,
		})
	}

	t.Render()
}
