package main

import (
	"log"

	"github.com/andreymgn/RSOI-poststats/pkg/poststats"
	"github.com/andreymgn/RSOI/pkg/tracer"
)

const (
	PoststatsAppID     = "PostStatsAPI"
	PoststatsAppSecret = "3BusyNfGQpyCr77J"
)

func runPostStats(port int, connString, jaegerAddr, redisAddr, redisPassword string, redisDB int) error {
	tracer, err := tracer.NewTracer("poststats", jaegerAddr)
	if err != nil {
		log.Fatal(err)
	}

	knownKeys := map[string]string{PoststatsAppID: PoststatsAppSecret}

	server, err := poststats.NewServer(connString, redisAddr, redisPassword, redisDB, knownKeys)
	if err != nil {
		log.Fatal(err)
	}

	return server.Start(port, tracer)
}
