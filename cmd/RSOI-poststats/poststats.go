package main

import (
	"github.com/andreymgn/RSOI-poststats/pkg/poststats"
)

func runPostStats(port int, connString string) error {
	server, err := poststats.NewServer(connString)
	if err != nil {
		return err
	}

	return server.Start(port)
}
