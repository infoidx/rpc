package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

var NewServer = grpc.NewServer

type Server struct {
	server *grpc.Server
	logger *grpclog.LoggerV2
}

func (s *Server) Serve(ctx context.Context) error {
	reflection.Register(s.server)
	// TODO
	return nil
}
