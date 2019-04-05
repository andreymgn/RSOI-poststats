package main

import (
	"github.com/andreymgn/RSOI-poststats/pkg/poststats"
	"github.com/andreymgn/RSOI/pkg/tracer"
)

func runPostStats(port int, connString, jaegerAddr string) error {
	tracer, closer, err := tracer.NewTracer("poststats", jaegerAddr)
	defer closer.Close()
	if err != nil {
		return err
	}

	server, err := poststats.NewServer(connString)
	if err != nil {
		return err
	}

	return server.Start(port, tracer)
}
