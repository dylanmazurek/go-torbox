package examples

import (
	"fmt"

	"github.com/dylanmazurek/go-torbox/pkg/torbox"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rs/zerolog/log"
)

func PrintUser(client *torbox.Client) {
	// Example: Get user information
	user, err := client.General.GetUser()
	if err != nil {
		log.Error().Err(err).Msg("failed to get user info")
		return
	}

	t := table.NewWriter()
	t.SetTitle("User Information")
	t.AppendHeader(table.Row{"Field", "Value"})
	t.AppendRow(table.Row{"ID", user.ID})
	t.AppendRow(table.Row{"Email", user.Email})
	t.AppendRow(table.Row{"Plan", user.Plan})
	dataDownloaded := user.TotalBytesDownloaded / 1024 / 1024 / 1024
	t.AppendRow(table.Row{"Data Downloaded", fmt.Sprintf("%d GB", dataDownloaded)})

	dataUploaded := user.TotalBytesUploaded / 1024 / 1024 / 1024
	t.AppendRow(table.Row{"Data Uploaded", fmt.Sprintf("%d GB", dataUploaded)})

	t.AppendRow(table.Row{"Torrents Downloaded", user.TorrentsDownloaded})
	t.AppendRow(table.Row{"Web Downloads Downloaded", user.WebDownloadsDownloaded})
	t.AppendRow(table.Row{"Usenet Downloads Downloaded", user.UsenetDownloadsDownloaded})

	t.AppendRow(table.Row{"Total Downloaded", user.TotalDownloaded})
	t.AppendRow(table.Row{"Total Uploaded", user.TotalUploaded})

	fmt.Println(t.Render())
}
