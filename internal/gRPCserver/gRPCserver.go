package grpcserver

import (
	"context"

	"github.com/matoous/go-nanoid/v2"
	"github.com/pushkinvladislav/link_shortening/api/shorter"
	"github.com/pushkinvladislav/link_shortening/config"
	"github.com/pushkinvladislav/link_shortening/internal/models"
	"github.com/pushkinvladislav/link_shortening/internal/repository/postgres"
	"github.com/pushkinvladislav/link_shortening/utils"
	"github.com/spf13/viper"
)

type GRPCServer struct{}

func (s *GRPCServer) Create(ctx context.Context, req *shorter.CreateRequest) (*shorter.CreateResponse, error) {

	config.Init()
	database := postgres.NewPostgres()

	_, err := database.EstablishPSQLConnection(&postgres.PSQlconfig{
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

	defer database.Close()

	URL := models.URL{
		LongURL:  req.GetLongURL(),
		ShortURL: "",
	}
	database.Link_shortening().FindLongURL(&URL)

	var shortURL string

	if URL.ShortURL == "" {
		shortURL, err = gonanoid.Generate("_0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 10)
		if err != nil {
			logger.Logger.Error("Failed to generate", err)
		}
		URL.ShortURL = shortURL

		_, err := database.Link_shortening().Create(&URL)
		if err != nil {
			logger.Logger.Error("Failed to added URL:", err)
			return nil, err
		}
	}
	shortURL = "https://localhost:8080/" + URL.ShortURL

	return &shorter.CreateResponse{ShortURL: shortURL}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *shorter.GetRequest) (*shorter.GetResponse, error) {

	config.Init()
	database := postgres.NewPostgres()

	_, err := database.EstablishPSQLConnection(&postgres.PSQlconfig{
		Host:     viper.GetString("db.postgres.host"),
		Port:     viper.GetString("db.postgres.port"),
		Password: viper.GetString("db.postgres.password"),
		DBName:   viper.GetString("db.postgres.database"),
		Username: viper.GetString("db.postgres.user"),
	})

	if err != nil {
		logger.Logger.Error(err)
	}

	URL := models.URL{
		LongURL:  "",
		ShortURL: req.GetShortURL(),
	}

	_, err1 := database.Link_shortening().Get(&URL)

	defer database.Close()
	return &shorter.GetResponse{LongURL: URL.LongURL}, err1
}
