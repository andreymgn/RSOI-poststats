package poststats

import (
	"fmt"
	"net"

	pb "github.com/andreymgn/RSOI-poststats/pkg/poststats/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server implements poststats service
type Server struct {
	db datastore
}

// NewServer returns a new server
func NewServer(connString string) (*Server, error) {
	db, err := newDB(connString)
	if err != nil {
		return nil, err
	}

	return &Server{db}, err
}

// Start starts a server
func (s *Server) Start(port int) error {
	creds, err := credentials.NewServerTLSFromFile("/cert.pem", "/key.pem")
	if err != nil {
		return err
	}

	server := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterPostStatsServer(server, s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return server.Serve(lis)
}
