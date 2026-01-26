package examples

import (
	"fmt"

	"github.com/aws/smithy-go/ptr"
	"github.com/dylanmazurek/go-torbox/pkg/torbox"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rs/zerolog/log"
)

func PrintActiveTorrents(client *torbox.Client) {
	torrents, err := client.General.GetActiveTorrents()
	if err != nil {
		log.Error().Err(err).Msg("failed to get active torrents")
		return
	}

	t := table.NewWriter()
	t.SetTitle("Active Torrents")
	t.AppendHeader(table.Row{"ID", "Name", "Progress", "Download Speed", "Upload Speed"})
	for _, torrent := range torrents {
		truncatedName := fmt.Sprintf("%.30s...", torrent.Name)
		t.AppendRow(table.Row{
			torrent.ID,
			truncatedName,
			fmt.Sprintf("%.2f%%", torrent.Progress*100),
			fmt.Sprintf("%d/s", torrent.DownloadSpeed),
			fmt.Sprintf("%d/s", torrent.UploadSpeed),
		})
	}

	fmt.Println(t.Render())
}

func PrintQueuedTorrents(client *torbox.Client) {
	queuedTorrents, err := client.General.GetQueuedTorrents(ptr.Bool(true))
	if err != nil {
		log.Error().Err(err).Msg("failed to get queued torrents")
		return
	}

	t := table.NewWriter()
	t.SetTitle("Queued Torrents")
	t.AppendHeader(table.Row{"ID", "Name"})
	for _, torrent := range queuedTorrents {
		truncatedName := fmt.Sprintf("%.30s...", torrent.Name)
		t.AppendRow(table.Row{
			torrent.ID,
			truncatedName,
		})
	}

	fmt.Println(t.Render())
}
