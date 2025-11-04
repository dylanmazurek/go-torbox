package main

import (
	"context"
	"os"

	"github.com/dylanmazurek/go-torbox/internal/logger"
	"github.com/dylanmazurek/go-torbox/pkg/torbox"

	"github.com/dylanmazurek/go-torbox/examples"
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

	examples.PrintUser(client)
	examples.PrintActiveTorrents(client)
	examples.PrintQueuedTorrents(client)
}
