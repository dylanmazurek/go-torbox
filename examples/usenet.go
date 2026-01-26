package examples

// func PrintActiveUsenet(client *torbox.Client) {
// 	usenetDownloads, err := client.General.GetActiveUsenetDownloads()
// 	if err != nil {
// 		log.Error().Err(err).Msg("failed to get active usenet downloads")
// 		return
// 	}

// 	t := table.NewWriter()
// 	t.SetTitle("Active Usenet Downloads")
// 	t.AppendHeader(table.Row{"ID", "Name", "Progress", "Download Speed"})
// 	for _, download := range usenetDownloads {
// 		truncatedName := fmt.Sprintf("%.30s...", download.Name)
// 		t.AppendRow(table.Row{
// 			download.ID,
// 			truncatedName,
// 			fmt.Sprintf("%.2f%%", download.Progress*100),
// 			fmt.Sprintf("%d/s", download.DownloadSpeed),
// 		})
// 	}

// 	fmt.Println(t.Render())
// }
