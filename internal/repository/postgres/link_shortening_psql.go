package postgres

import (
	"context"
	"github.com/pushkinvladislav/link_shortening/internal/models"
)

type Link_shortening struct {
	postgres *Postgres
}

func (r *Link_shortening) Create(l *models.URL) (*models.URL, error) {

	CreateQuery := "INSERT INTO URLs (longURL, shortURL) values ($1, $2) RETURNING id;"

	if err := r.postgres.db.QueryRow(context.Background(), CreateQuery, l.LongURL, l.ShortURL).Scan(&l.Id); err != nil {
		return nil, err
	}
	return l, nil
}

func (r *Link_shortening) Get(l *models.URL) (*models.URL, error) {

	GetQuery := "SELECT longURL FROM URLs WHERE shortURL=$1"

	if err := r.postgres.db.QueryRow(context.Background(), GetQuery, l.ShortURL).Scan(&l.LongURL); err != nil {
		return nil, err
	}
	return l, nil
}

func (r *Link_shortening) FindLongURL(l *models.URL) (*models.URL, error) {

	FindLongURLQuery := "SELECT shortURL from URLs WHERE longURL=$1"

	if err := r.postgres.db.QueryRow(context.Background(), FindLongURLQuery, l.LongURL).Scan(&l.ShortURL); err != nil {
		return nil, err
	}
	return l, nil

}
