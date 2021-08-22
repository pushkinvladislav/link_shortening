package grpcserver

import (
	"context"
	"github.com/pushkinvladislav/link_shortening/api/shorter"
	"github.com/pushkinvladislav/link_shortening/utils"
	"github.com/matoous/go-nanoid/v2"
)

type GRPCServer struct{}

func (s *GRPCServer) Create(ctx context.Context, req *shorter.CreateRequest) (*shorter.CreateResponse, error) {
	
	shortURL, err := gonanoid.Generate("_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 10)
		if err != nil {
			logger.Logger.Error("Failed to generate",err)
		}
	return &shorter.CreateResponse{ShortURL: shortURL}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *shorter.GetRequest) (*shorter.GetResponse, error) {
	return &shorter.GetResponse{LongURL: "kmkjk"}, nil
}
