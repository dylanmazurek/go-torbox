package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dylanmazurek/go-torbox/pkg/torbox"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("TORBOX_API_KEY")
	if apiKey == "" {
		log.Fatal("TORBOX_API_KEY environment variable not set")
	}

	// Create client
	client, err := torbox.New(ctx, torbox.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example 1: Check if torrents are cached
	fmt.Println("=== Checking Cached Torrents ===")
	hashes := []string{"example_hash_1", "example_hash_2"}
	cachedResp, err := client.General.CheckCachedTorrents(hashes)
	if err != nil {
		log.Printf("Error checking cached: %v", err)
	} else {
		fmt.Printf("Checked %d hashes, success: %t\n", len(hashes), cachedResp.Success)
	}

	// Example 2: Get usage statistics
	fmt.Println("\n=== Getting Statistics ===")
	stats, err := client.General.GetStats()
	if err != nil {
		log.Printf("Error getting stats: %v", err)
	} else {
		fmt.Printf("Total Downloaded: %d bytes\n", stats.TotalDownloaded)
		fmt.Printf("Total Uploaded: %d bytes\n", stats.TotalUploaded)
		fmt.Printf("Active Torrents: %d\n", stats.ActiveTorrents)
		fmt.Printf("Queued Torrents: %d\n", stats.QueuedTorrents)
	}

	// Example 3: Get notifications
	fmt.Println("\n=== Getting Notifications ===")
	notifications, err := client.General.GetNotifications()
	if err != nil {
		log.Printf("Error getting notifications: %v", err)
	} else {
		fmt.Printf("Found %d notifications\n", len(notifications))
		for _, notif := range notifications {
			fmt.Printf("- [%s] %s: %s\n", notif.Type, notif.Title, notif.Message)
		}
	}

	// Example 4: Get usenet downloads
	fmt.Println("\n=== Getting Usenet Downloads ===")
	usenetDownloads, err := client.General.GetUsenetList(nil)
	if err != nil {
		log.Printf("Error getting usenet downloads: %v", err)
	} else {
		fmt.Printf("Found %d usenet downloads\n", len(usenetDownloads))
		for _, dl := range usenetDownloads {
			fmt.Printf("- %s (%.2f%% complete)\n", dl.Name, dl.Progress)
		}
	}

	// Example 5: Get web downloads
	fmt.Println("\n=== Getting Web Downloads ===")
	webDownloads, err := client.General.GetWebDownloadList(nil)
	if err != nil {
		log.Printf("Error getting web downloads: %v", err)
	} else {
		fmt.Printf("Found %d web downloads\n", len(webDownloads))
		for _, dl := range webDownloads {
			fmt.Printf("- %s (%.2f%% complete)\n", dl.Name, dl.Progress)
		}
	}

	// Example 6: Get RSS feeds
	fmt.Println("\n=== Getting RSS Feeds ===")
	rssFeeds, err := client.General.GetRSSFeeds()
	if err != nil {
		log.Printf("Error getting RSS feeds: %v", err)
	} else {
		fmt.Printf("Found %d RSS feeds\n", len(rssFeeds))
		for _, feed := range rssFeeds {
			status := "inactive"
			if feed.Active {
				status = "active"
			}
			fmt.Printf("- %s (%s): %s\n", feed.Name, status, feed.URL)
		}
	}

	// Example 7: Convert magnet to torrent file
	fmt.Println("\n=== Converting Magnet to File ===")
	magnetLink := "magnet:?xt=urn:btih:example"
	torrentFile, err := client.General.MagnetToFile(magnetLink)
	if err != nil {
		log.Printf("Error converting magnet: %v", err)
	} else {
		fmt.Printf("Torrent file created (length: %d bytes)\n", len(*torrentFile))
	}

	// Example 8: Export torrent data
	fmt.Println("\n=== Exporting Torrent Data ===")
	activeTorrents, err := client.General.ListActiveTorrents(false)
	if err != nil {
		log.Printf("Error listing torrents: %v", err)
	} else if len(activeTorrents) > 0 {
		exportData, err := client.General.ExportTorrentData(activeTorrents[0].ID, "json")
		if err != nil {
			log.Printf("Error exporting data: %v", err)
		} else {
			fmt.Printf("Exported data (length: %d bytes)\n", len(*exportData))
		}
	}

	fmt.Println("\n=== Examples Complete ===")
}
