package postgres

import (
	"fmt"
	"context"
	"github.com/pushkinvladislav/link_shortening/utils"
	"github.com/jackc/pgx/v4"
)

type Postgres struct {
	db             *pgx.Conn
	linkshorteningPSQL *Link_shortening
	
	
}
type PSQlconfig struct{
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}


func (s *Postgres) EstablishPSQLConnection(cnf *PSQlconfig) (*pgx.Conn, error) {

	db, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.DBName))
	if err != nil {
		logger.Logger.Error(err)
	}
	s.db = db

	return s.db, nil
}

func (s *Postgres) Close() {
	s.db.Close(context.Background())
}

func (s *Postgres) Link_shortening() *Link_shortening {
	if s.linkshorteningPSQL != nil {
		return s.linkshorteningPSQL
	}

	s.linkshorteningPSQL = &Link_shortening{
		postgres: s,
	}

	return s.linkshorteningPSQL

}

func NewPostgres() *Postgres {
	return &Postgres{}
}
