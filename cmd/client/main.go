package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pushkinvladislav/link_shortening/api/shorter"
	"github.com/pushkinvladislav/link_shortening/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	logger.InitLogger()
	logger.Logger.Info("Starting client...")


    args := os.Args
    conn, err := grpc.Dial("8080",grpc.WithInsecure())

    if err != nil {
        grpclog.Fatalf("fail to dial: %v", err)
    }

    defer conn.Close()

	client := shorter.NewShorterServiceClient(conn)
    request := &shorter.CreateRequest{
        LongURL: args[1],
    }
    response, err := client.Create(context.Background(), request)

    if err != nil {
        grpclog.Fatalf("fail to dial: %v", err)
    }

   fmt.Println(response.ShortURL)
}
