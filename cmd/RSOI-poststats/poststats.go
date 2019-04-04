package main

import (
	"log"

	"github.com/andreymgn/RSOI-poststats/pkg/poststats"
)

func runPostStats(port int, connString string) error {
	server, err := poststats.NewServer(connString)
	if err != nil {
		log.Fatal(err)
	}

	return server.Start(port)
}
