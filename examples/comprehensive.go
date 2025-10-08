package main

import (
	"context"
	"fmt"
	"os"

	"github.com/dylanmazurek/go-torbox/internal/logger"
	"github.com/dylanmazurek/go-torbox/pkg/torbox"
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

	// Example: Get user information
	fmt.Println("=== User Information ===")
	user, err := client.General.GetUser()
	if err != nil {
		log.Error().Err(err).Msg("failed to get user info")
	} else {
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Printf("Plan: %s\n", user.Plan)
		fmt.Printf("Total Downloaded: %d bytes\n", user.TotalDownloaded)
		fmt.Printf("Total Uploaded: %d bytes\n", user.TotalUploaded)
	}

	// Example: Get account statistics
	fmt.Println("\n=== Account Statistics ===")
	stats, err := client.General.GetStats()
	if err != nil {
		log.Error().Err(err).Msg("failed to get stats")
	} else {
		fmt.Printf("Total Torrents: %d\n", stats.TotalTorrents)
		fmt.Printf("Active Torrents: %d\n", stats.ActiveTorrents)
		fmt.Printf("Queued Torrents: %d\n", stats.QueuedTorrents)
		fmt.Printf("Total Usenet: %d\n", stats.TotalUsenet)
		fmt.Printf("Total Web DL: %d\n", stats.TotalWebDL)
		fmt.Printf("Used Space: %d bytes\n", stats.UsedSpace)
		fmt.Printf("Available Space: %d bytes\n", stats.AvailableSpace)
	}

	// Example: List active torrents
	fmt.Println("\n=== Active Torrents ===")
	activeTorrents, err := client.General.GetActiveTorrents()
	if err != nil {
		log.Error().Err(err).Msg("failed to get active torrents")
	} else {
		if len(activeTorrents) == 0 {
			fmt.Println("No active torrents")
		} else {
			for i, t := range activeTorrents {
				fmt.Printf("%d. %s (%.2f%%) - %s\n", i+1, t.Name, t.Progress*100, t.DownloadState)
			}
		}
	}

	// Example: List queued torrents
	fmt.Println("\n=== Queued Torrents ===")
	queuedTorrents, err := client.General.GetQueuedTorrents()
	if err != nil {
		log.Error().Err(err).Msg("failed to get queued torrents")
	} else {
		if len(queuedTorrents) == 0 {
			fmt.Println("No queued torrents")
		} else {
			for i, t := range queuedTorrents {
				fmt.Printf("%d. %s (Hash: %s)\n", i+1, t.Name, t.Hash)
			}
		}
	}

	// Example: List usenet downloads
	fmt.Println("\n=== Usenet Downloads ===")
	usenetList, err := client.General.GetUsenetList()
	if err != nil {
		log.Error().Err(err).Msg("failed to get usenet list")
	} else {
		if len(usenetList) == 0 {
			fmt.Println("No usenet downloads")
		} else {
			for i, u := range usenetList {
				fmt.Printf("%d. %s (%.2f%%) - %s\n", i+1, u.Name, u.Progress*100, u.DownloadState)
			}
		}
	}

	// Example: Get notifications
	fmt.Println("\n=== Notifications ===")
	notifications, err := client.General.GetNotifications()
	if err != nil {
		log.Error().Err(err).Msg("failed to get notifications")
	} else {
		if len(notifications) == 0 {
			fmt.Println("No notifications")
		} else {
			for i, n := range notifications {
				readStatus := "Unread"
				if n.Read {
					readStatus = "Read"
				}
				fmt.Printf("%d. [%s] %s: %s\n", i+1, readStatus, n.Title, n.Message)
			}
		}
	}

	// Example: Get integration jobs
	fmt.Println("\n=== Integration Jobs ===")
	jobs, err := client.General.GetIntegrationJobs()
	if err != nil {
		log.Error().Err(err).Msg("failed to get integration jobs")
	} else {
		if len(jobs) == 0 {
			fmt.Println("No integration jobs")
		} else {
			for i, j := range jobs {
				fmt.Printf("%d. %s -> %s (%.2f%%) - %s\n", i+1, j.FileName, j.Destination, j.Progress*100, j.Status)
			}
		}
	}

	// Example: Check if a torrent is cached (using a sample hash)
	// Note: Replace with an actual torrent hash to test
	sampleHash := "abc123def456"
	fmt.Printf("\n=== Cache Check for Hash: %s ===\n", sampleHash)
	cacheInfo, err := client.General.CheckCached(sampleHash)
	if err != nil {
		log.Error().Err(err).Msg("failed to check cache")
	} else {
		if cacheInfo.Cached {
			fmt.Printf("Cached! Name: %s, Size: %d bytes\n", cacheInfo.Name, cacheInfo.Size)
		} else {
			fmt.Println("Not cached")
		}
	}

	// Example: Search torrents
	fmt.Println("\n=== Search Torrents (query: ubuntu) ===")
	searchResults, err := client.General.SearchTorrents("ubuntu")
	if err != nil {
		log.Error().Err(err).Msg("failed to search torrents")
	} else {
		if len(searchResults) == 0 {
			fmt.Println("No results found")
		} else {
			for i, t := range searchResults[:min(5, len(searchResults))] {
				fmt.Printf("%d. %s (Size: %d bytes)\n", i+1, t.Name, t.Size)
			}
		}
	}

	// Example: Create a web download (commented out to prevent accidental creation)
	/*
	fmt.Println("\n=== Creating Web Download ===")
	link := "https://example.com/file.zip"
	name := "Example File"
	webReq := models.CreateWebDownloadRequest{
		Link: link,
		Name: &name,
	}
	webDL, err := client.General.CreateWebDownload(webReq)
	if err != nil {
		log.Error().Err(err).Msg("failed to create web download")
	} else {
		fmt.Printf("Created web download: %s (ID: %d)\n", webDL.Name, webDL.ID)
	}
	*/

	fmt.Println("\n=== Demo Complete ===")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
