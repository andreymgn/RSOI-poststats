package main

import (
	"log"

	"github.com/andreymgn/RSOI-poststats/pkg/poststats"
)

const (
	PoststatsAppID     = "PostStatsAPI"
	PoststatsAppSecret = "3BusyNfGQpyCr77J"
)

func runPostStats(port int, connString, redisAddr, redisPassword string, redisDB int) error {
	knownKeys := map[string]string{PoststatsAppID: PoststatsAppSecret}

	server, err := poststats.NewServer(connString, redisAddr, redisPassword, redisDB, knownKeys)
	if err != nil {
		log.Fatal(err)
	}

	return server.Start(port)
}
