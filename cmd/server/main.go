package main

import (
	"net"
	"time"

	"github.com/pushkinvladislav/link_shortening/api/shorter"
	"github.com/pushkinvladislav/link_shortening/internal/gRPCserver"
	"github.com/pushkinvladislav/link_shortening/utils"
	"google.golang.org/grpc"
)

func main() {
	logger.InitLogger()
	logger.Logger.Info("Starting server...")
	time.Sleep(time.Second * 4)

	s := grpc.NewServer()
	srv := &grpcserver.GRPCServer{}
	shorter.RegisterShorterServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Logger.Error(err)
	} else {
		logger.Logger.Info("Server listening on :8080")
	}

	if err := s.Serve(l); err != nil {
		logger.Logger.Error(err)
	}

}
