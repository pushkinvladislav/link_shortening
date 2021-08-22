package main

import (
	"net"
	"time"

	"github.com/pushkinvladislav/link_shortening/api/shorter"
	"github.com/pushkinvladislav/link_shortening/config"
	"github.com/pushkinvladislav/link_shortening/internal/repository/postgres"
	"github.com/pushkinvladislav/link_shortening/internal/gRPCserver"
	"github.com/pushkinvladislav/link_shortening/utils"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	logger.InitLogger()
	logger.Logger.Info("Starting server...")
	time.Sleep(time.Second * 4)

	config.Init()

	logger.Logger.Info("connecting to the database...")
	conn, err := postgres.EstablishPSQLConnection(&postgres.PSQLConfig{
		Host:     viper.GetString("db.postgres.host"),
		Port:     viper.GetString("db.postgres.port"),
		Password: viper.GetString("db.postgres.password"),
		DBName:   viper.GetString("db.postgres.database"),
		Username: viper.GetString("db.postgres.user"),
		SSLMode:  viper.GetString("db.postgres.sslmode"),
	})
	if err != nil {
		logger.Logger.Error(err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			logger.Logger.Error(err.Error())
		}
	}()
	logger.Logger.Info("postgres connection established")

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
